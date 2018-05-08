package singleton

import (
	"fmt"
	"sync"
)

//https://blog.csdn.net/qibin0506/article/details/50733314

var m *Manager

func GetInstance1() *Manager {
	if m == nil {
		m = &Manager {}
	}
	return m
}

type Manager struct {}

func (p Manager) Manage() {
	fmt.Println("manage...")
}

/////////////////////

var lock *sync.Mutex = &sync.Mutex {}

func GetInstance2() *Manager {
	lock.Lock()
	defer lock.Unlock()
	if m == nil {
		m = &Manager {}
	}
	return m
}

/////////////////////

func GetInstance3() *Manager {
	if m == nil {
		lock.Lock()
		defer lock.Unlock()
		if m == nil {
			m = &Manager {}
		}
	}

	return m
}

/////////////////////////

var once sync.Once

func GetInstance() *Manager {
	once.Do(func() {
		m = &Manager {}
	})
	return m
}
