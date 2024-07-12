package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type Stub struct {
	Name     string       `json:"name"`
	Request  StubRequest  `json:"request"`
	Response StubResponse `json:"response"`
}

type StubRequest struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Body    string            `json:"body"`
	Cookie  string            `json:"cookie"`
	Headers map[string]string `json:"headers"`
}

type StubResponse struct {
	Status  int               `json:"status"`
	Body    string            `json:"body"`
	Headers map[string]string `json:"headers"`
}

type Stubs struct {
	Map   map[string]*Stub
	Mutex sync.RWMutex
}

func NewStubs() *Stubs {
	return &Stubs{
		Map: make(map[string]*Stub),
	}
}

func (stubs *Stubs) Add(stub *Stub) {
	// Стоит ли сделать проверку на обратное преобразование в json строку?
	// Или же таскать его за собой, но это избыток данных
	request, _ := json.Marshal(stub.Request)
	requestString := string(request)
	stubs.Mutex.Lock()
	defer stubs.Mutex.Unlock()
	stubs.Map[requestString] = stub
}

func (stubs *Stubs) Get(requestString *string) *Stub {
	stubs.Mutex.RLock()
	defer stubs.Mutex.RUnlock()
	stub, ok := stubs.Map[*requestString]
	if !ok {
		return nil
	}
	return stub
}

// TODO Проверка на валидность порта
// Сгруппировать функции в main (а оно надо?)
// GET для генератора заглушек
// Журналирование
// Добавление задержки, критических ошибок по типу отключения сокета
// Добавление regex
// Разбить по модулям
// Можно добавить флаг для другого URL вместо generate
// Рассмотреть случай, если порт занят
// Вынести строку generate в константы
// Проверка на ошибки

func main() {
	port := flag.String("port", "8080", "port to listen on")
	flag.Parse()
	stubs := NewStubs()
	http.Handle("/generate", &generateHandler{stubsPtr: stubs})
	http.Handle("/", &stubHandler{stubsPtr: stubs})

	fmt.Printf("Server listening on port %s", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

type generateHandler struct {
	stubsPtr *Stubs
}

func (h *generateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// TODO валидация данных запроса
	// Стоит отделить указатели от полноценных записей припиской Ptr
	stub := new(Stub)
	err := json.NewDecoder(r.Body).Decode(stub)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	h.stubsPtr.Add(stub)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Mock created successfully for %s on %s\n", stub.Name,
		stub.Request.URL)
}

type stubHandler struct {
	stubsPtr *Stubs
}

func parseRequestBody(r *http.Request) (*string, error) {
	body, err := (io.ReadAll(r.Body))
	if err != nil {
		return nil, err
	}

	r.Body = io.NopCloser(bytes.NewReader(body))

	bodyString := string(body)

	return &bodyString, nil
}

// TODO Валидация
func parseRequestString(r *http.Request) (*string, error) {
	request := StubRequest{
		Method: r.Method,
		URL:    r.URL.String(),
		//Headers: make(map[string]string),
	}
	// Гибкая настройка того, что нужно игнорировать
	//
	//for key, values := range r.Header {
	//	request.Headers[key] = values[0]
	//}
	//
	//    for _, cookie := range r.Cookies() {
	//        request.Cookie += cookie.String() + "; "
	//    }

	requestBodyPtr, err := parseRequestBody(r)
	if err != nil {
		return nil, err
	}
	request.Body = *requestBodyPtr

	requestByte, _ := json.Marshal(request)
	requestString := string(requestByte)

	return &requestString, nil
}

func (h *stubHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//if r.Method != http.MethodGet {
	//	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	//	return
	//}

	// Наименование функций бы поменять
	requestString, err := parseRequestString(r)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	stub := h.stubsPtr.Get(requestString)
	if stub == nil {
		http.Error(w, "Stub not found", http.StatusNotFound)
		return
	}
	// Перепроверить, соотв. ли статус ошибке

	w.WriteHeader(stub.Response.Status)
	for key, value := range stub.Response.Headers {
		w.Header().Set(key, value)
	}

	fmt.Fprintf(w, "%s", stub.Response.Body)
}
