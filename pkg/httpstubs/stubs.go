package httpstubs

import (
	"net/http"
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
		if stub.Name != nil && *item.Name == *stub.Name {
			stubs.Items[index] = stub
			return index
		}
	}
	stubs.Items = append(stubs.Items, stub)
	stubs.IdMap[len(stubs.IdMap)+1] = stub
	return len(stubs.Items) - 1
}

func (stubs *Stubs) GetMatchingStubsByRequest(r *http.Request) []*Stub {
	stubs.Mutex.RLock()
	defer stubs.Mutex.RUnlock()

	matchingStubs := make([]*Stub, 0, 10)

	for _, stub := range stubs.Items {
		if stub.Request.Matches(r) {
			matchingStubs = append(matchingStubs, stub)
		}
	}

	return matchingStubs
}

func (stubs *Stubs) GetById(id int) *Stub {
	stubs.Mutex.RLock()
	defer stubs.Mutex.RUnlock()

	return stubs.IdMap[id]
}
