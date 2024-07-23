package httpstubs

import (
	"sync"
)

// Хотелось бы использовать хэш, однако сопоставление настоящего запроса к заглушкам
// не предполагает связи один к одному и является комплексной, потому имплементация в виде простого слайса

// Изменен на map, где ключ это id, будут накладки, зато удобно

type Stubs struct {
	Map   map[int]*Stub
	Mutex sync.RWMutex
}

func NewStubs() *Stubs {
	return &Stubs{
		Map: make(map[int]*Stub, 10),
	}
}

func (stubs *Stubs) Add(stub *Stub) int {
	stubs.Mutex.Lock()
	defer stubs.Mutex.Unlock()

	for id, stubItem := range stubs.Map {
		if stub.Name != nil && *stubItem.Name == *stub.Name {
			stubs.Map[id] = stub
			return id
		}
	}
	newId := len(stubs.Map) + 1
	stubs.Map[newId] = stub
	return newId
}

func (stubs *Stubs) Put(stub *Stub, id int) bool {
	stubs.Mutex.Lock()
	defer stubs.Mutex.Unlock()

	if _, ok := stubs.Map[id]; ok {
		stubs.Map[id] = stub
		return true
	}
	return false
}

func (stubs *Stubs) Delete(id int) bool {
	stubs.Mutex.Lock()
	defer stubs.Mutex.Unlock()

	if _, ok := stubs.Map[id]; ok {
		delete(stubs.Map, id)
		return true
	}
	return false
}

func (stubs *Stubs) GetById(id int) *Stub {
	stubs.Mutex.RLock()
	defer stubs.Mutex.RUnlock()

	return stubs.Map[id]
}

func (stubs *Stubs) GetMatchingStubsByRequest(r *MyRequest) []*Stub {
	stubs.Mutex.RLock()
	defer stubs.Mutex.RUnlock()

	matchingStubs := make([]*Stub, 0, 10)

	for _, stub := range stubs.Map {
		if stub.Matches(r) {
			matchingStubs = append(matchingStubs, stub)
		}
	}

	return matchingStubs
}

func (stubs *Stubs) OptimizedGetMatchingStubsByRequest(r *MyRequest) []*Stub {
	stubs.Mutex.RLock()
	defer stubs.Mutex.RUnlock()

	var priority int
	isFoundOnce := false
	matchingStubs := make([]*Stub, 0, 10)

	for _, stub := range stubs.Map {
		if !isFoundOnce {
			if stub.Matches(r) {
				matchingStubs = append(matchingStubs, stub)
				isFoundOnce = true
				priority = stub.Properties.Priority
			}
		} else {
			if stub.Properties.Priority > priority {
				if stub.Matches(r) {
					matchingStubs = []*Stub{stub}
					priority = stub.Properties.Priority
				}
			} else if stub.Properties.Priority == priority {
				if stub.Matches(r) {
					matchingStubs = append(matchingStubs, stub)
					priority = stub.Properties.Priority
				}
			}
		}
	}

	return matchingStubs
}
