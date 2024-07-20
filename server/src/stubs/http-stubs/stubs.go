package httpstubs

import (
	"bytes"
	"fmt"
	"regexp"
	"sync"
)

// Хотелось бы использовать что-то по типу map[request]response, однако сопоставление настоящего запроса к заглушкам
// не предполагает связи один к одному и является комплексной, потому имплементация в виде простого слайса

type Stubs struct {
	Items []*Stub
	// Можно было бы обойтись одним, но не знаю, насколько медленее хэш по индексу работает медленее простой индексации
	IdMap map[int]*Stub
	Mutex sync.RWMutex
}

func NewStubs() *Stubs {
	return &Stubs{
		Items: make([]*Stub, 0, 10),
		IdMap: make(map[int]*Stub),
	}
}

func (stubs *Stubs) Add(stub *Stub) int {
	stubs.Mutex.Lock()
	defer stubs.Mutex.Unlock()

	// Для удобства тестирования; если имя заглушки совпадает, то переписывает
	// Ломает REST, ну и ладно, зато удобно
	for index, item := range stubs.Items {
		if item.Name == stub.Name {
			stubs.Items[index] = stub
			return index
		}
	}
	stubs.Items = append(stubs.Items, stub)
	stubs.IdMap[len(stubs.IdMap)+1] = stub
	return len(stubs.Items) - 1
}

func (stubs *Stubs) GetMatchingStubsByRequest(request *StubRequest) []*Stub {
	stubs.Mutex.RLock()
	defer stubs.Mutex.RUnlock()

	matchingStubs := make([]*Stub, 0, 10)

	for _, stub := range stubs.Items {
		if doRequestsMatch(request, &stub.Request) {
			matchingStubs = append(matchingStubs, stub)
		}
	}

	return matchingStubs
}

func doRequestsMatch(request *StubRequest, stubRequest *StubRequest) bool {
	return doURLsMatch(request.URL, stubRequest.URL, stubRequest.URLMatches) &&
		doMethodsMatch(request.Method, stubRequest.Method) &&
		doBodiesMatch(request.BodyBinary, stubRequest.BodyBinary) &&
		doHeadersMatch(request.Headers, stubRequest.Headers)
}

func doURLsMatch(requestURL string, stubURL string, stubURLMatches string) bool {
	if stubURLMatches != "" {
		re, err := regexp.Compile(stubURLMatches)
		if err != nil {
			return false
		}
		fmt.Println(re.MatchString(requestURL))
		return re.MatchString(requestURL)

	}

	return requestURL == stubURL
}

func doMethodsMatch(requestMethod string, stubMethod string) bool {
	return requestMethod == stubMethod
}

func doBodiesMatch(requestBody []byte, stubBody []byte) bool {
	return stubBody == nil || bytes.Equal(stubBody, requestBody)
}

func doHeadersMatch(requestHeaders map[string]string, stubHeaders map[string]string) bool {
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

func (stubs *Stubs) GetById(id int) *Stub {
	stubs.Mutex.RLock()
	defer stubs.Mutex.RUnlock()

	return stubs.IdMap[id]
}
