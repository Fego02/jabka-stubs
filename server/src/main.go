package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type Stub struct {
	Name       string         `json:"name"`
	Request    StubRequest    `json:"request"`
	Response   StubResponse   `json:"response"`
	Properties StubProperties `json:"properties"`
}

type StubRequest struct {
	Method     string            `json:"method"`
	URL        string            `json:"url"`
	Body       string            `json:"body"`
	Headers    map[string]string `json:"headers"`
	BodyBinary []byte
}

type StubResponse struct {
	Status     int               `json:"status"`
	Body       string            `json:"body"`
	Headers    map[string]string `json:"headers"`
	BodyBinary []byte
}

type StubProperties struct {
	IsLoggingEnabled bool `json:"is_logging_enabled"`
	Delay            int  `json:"delay"`
}

type Stubs struct {
	Items []*Stub
	Mutex sync.RWMutex
}

func NewStubs() *Stubs {
	return &Stubs{
		Items: make([]*Stub, 0, 10),
	}
}

func (stubs *Stubs) Add(stub *Stub) {
	stubs.Mutex.Lock()
	defer stubs.Mutex.Unlock()

	for index, item := range stubs.Items {
		if item.Name == stub.Name {
			stubs.Items[index] = stub
			return
		}
	}
	stubs.Items = append(stubs.Items, stub)
}

func (stubs *Stubs) Get(stubRequest *StubRequest) []*Stub {
	stubs.Mutex.RLock()
	defer stubs.Mutex.RUnlock()

	resultSlice := make([]*Stub, 0, 10)

	for _, stub := range stubs.Items {
		if stub.Request.URL == stubRequest.URL && stub.Request.Method == stubRequest.Method {
			if stub.Request.BodyBinary == nil || bytes.Equal(stub.Request.BodyBinary, stubRequest.BodyBinary) {
				if areHeadersMatched(stubRequest.Headers, stub.Request.Headers) {
					resultSlice = append(resultSlice, stub)
				}
			}
		}
	}

	return resultSlice
}

func areHeadersMatched(requestHeaders map[string]string, stubHeaders map[string]string) bool {
	for stubHeaderKey, stubHeaderValue := range stubHeaders {
		requestHeaderValue, ok := requestHeaders[stubHeaderKey]
		if !ok {
			return false
		}

		if stubHeaderValue != requestHeaderValue {
			return false
		}
	}
	return true
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
	address := flag.String("address", "127.0.0.1", "address to bind to")
	flag.Parse()
	stubs := NewStubs()
	http.Handle("/generate", &generateHandler{stubsPtr: stubs})
	http.Handle("/", &stubHandler{stubsPtr: stubs})

	fmt.Printf("Server listening on %s:%s", *address, *port)
	log.Fatal(http.ListenAndServe(*address+":"+*port, nil))
}

type generateHandler struct {
	stubsPtr *Stubs
}

func (h *generateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	contentType := r.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "multipart/form-data") && contentType != "application/json" {
		http.Error(w, "Invalid Content Type", http.StatusBadRequest)
		return
	}
	var err error
	var stubData io.Reader = r.Body
	// TODO валидация данных запроса
	// Стоит отделить указатели от полноценных записей припиской Ptr
	stub := new(Stub)
	stub.Request.Body = "NULL"
	stub.Response.Body = "NULL"

	if strings.HasPrefix(contentType, "multipart/form-data") {
		err = r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Invalid multipart", http.StatusBadRequest)
			return
		}
		stubDataFileHeader, ok := r.MultipartForm.File["stub-data"]

		if !ok || len(stubDataFileHeader) == 0 {
			http.Error(w, "Stub data not found", http.StatusBadRequest)
			return
		}

		stubDataFile, err := stubDataFileHeader[0].Open()
		if err != nil {
			http.Error(w, "Cannon open stub data", http.StatusBadRequest)
			return
		}
		defer stubDataFile.Close()
		stubData = stubDataFile

	}

	err = json.NewDecoder(stubData).Decode(stub)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if stub.Name == "" {
		http.Error(w, "Invalid Name", http.StatusBadRequest)
		return
	}

	_, err = url.ParseRequestURI(stub.Request.URL)
	if err != nil {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	switch stub.Request.Method {
	case http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch,
		http.MethodDelete, http.MethodConnect, http.MethodOptions, http.MethodTrace:
	default:
		http.Error(w, "Invalid Method", http.StatusBadRequest)
		return
	}

	if stub.Response.Status < 100 || stub.Response.Status > 999 {
		http.Error(w, "Invalid Status Code", http.StatusBadRequest)
		return
	}

	for headerKey := range stub.Request.Headers {
		if headerKey == "" {
			http.Error(w, "Invalid Request Headers Data", http.StatusBadRequest)
			return
		}
	}
	for headerKey := range stub.Response.Headers {
		if headerKey == "" {
			http.Error(w, "Invalid Response Headers Data", http.StatusBadRequest)
			return
		}
	}

	if strings.HasPrefix(contentType, "multipart/form-data") {
		requestBodyFileHeader, ok := r.MultipartForm.File["request-body"]
		if !ok || len(requestBodyFileHeader) == 0 {
			if stub.Request.Body != "NULL" {
				stub.Request.BodyBinary = []byte(stub.Request.Body)
			}
		}
		requestBodyFile, err := requestBodyFileHeader[0].Open()
		if err != nil {
			http.Error(w, "Cannon open request body file", http.StatusBadRequest)
			return
		}
		defer requestBodyFile.Close()

		stub.Request.BodyBinary, err = io.ReadAll(requestBodyFile)
		if err != nil {
			http.Error(w, "Cannon read request body file", http.StatusBadRequest)
			return
		}

		responseBodyFileHeader, ok := r.MultipartForm.File["response-body"]
		if !ok || len(responseBodyFileHeader) == 0 {
			if stub.Response.Body != "NULL" {
				stub.Response.BodyBinary = []byte(stub.Response.Body)
			}
		}
		responseBodyFile, err := responseBodyFileHeader[0].Open()
		if err != nil {
			http.Error(w, "Cannon open response body file", http.StatusBadRequest)
			return
		}
		defer responseBodyFile.Close()

		stub.Response.BodyBinary, err = io.ReadAll(responseBodyFile)
		if err != nil {
			http.Error(w, "Cannon read response body file", http.StatusBadRequest)
			return
		}

	} else {
		if stub.Request.Body != "NULL" {
			stub.Request.BodyBinary = []byte(stub.Request.Body)
		}
		if stub.Response.Body != "NULL" {
			stub.Response.BodyBinary = []byte(stub.Response.Body)
		}
	}

	h.stubsPtr.Add(stub)

	w.WriteHeader(http.StatusCreated)
	_, err = fmt.Fprintf(w, "Stub created successfully for %s on %s\n", stub.Name, stub.Request.URL)
	if err != nil {
		http.Error(w, "Writing Error", http.StatusInternalServerError)
	}
}

type stubHandler struct {
	stubsPtr *Stubs
}

func readRequestBody(r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close() // Важно закрыть тело запроса после чтения

	return body, nil
}

// TODO Валидация
func readRequest(r *http.Request) (*StubRequest, error) {
	request := new(StubRequest)
	request.Method = r.Method
	request.URL = r.URL.String()
	request.Headers = make(map[string]string)

	for key, values := range r.Header {
		request.Headers[key] = values[0]
	}

	body, err := readRequestBody(r)
	if err != nil {
		return nil, err
	}

	request.BodyBinary = body

	return request, nil
}

func (h *stubHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Наименование функций бы поменять
	request, err := readRequest(r)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	stubs := h.stubsPtr.Get(request)

	if len(stubs) == 0 {
		http.Error(w, "Stub not found", http.StatusNotFound)
		return
	}

	if len(stubs) != 1 {
		http.Error(w, "Multiple stubs found", http.StatusNotFound)
		return
	}

	stub := stubs[0]

	w.WriteHeader(stub.Response.Status)
	for key, value := range stub.Response.Headers {
		w.Header().Set(key, value)
	}

	_, err = w.Write(stub.Response.BodyBinary)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Lol")
		http.Error(w, "Writing Error", http.StatusInternalServerError)
	}
}
