package hashmap

import (
	"sync"
	"fmt"
)

type HashNode interface {
	Key() string
}

type HashTable struct {
	sync.Mutex
	hashMap map[string]HashNode
}

func (t *HashTable)Add(n HashNode) error {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()

	if t == nil {
		return fmt.Errorf("hashtable is nil")
	}

	_, ok := t.hashMap[n.Key()]
	if ok {
		return fmt.Errorf("add duplicated node")
	}
	t.hashMap[n.Key()] = n

	return nil
}

func (t *HashTable)Del(n HashNode) error {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()

	if t == nil {
		return fmt.Errorf("hashtable is nil")
	}
	_, ok := t.hashMap[n.Key()]
	if !ok {
		return fmt.Errorf("cannot find node")
	}
	delete(t.hashMap, n.Key())
	return nil
}
