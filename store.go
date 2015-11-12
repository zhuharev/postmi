package postmi

import (
	"fmt"
	"sync"
)

type Store interface {
	Connect(string) error

	Save(*Post) error
	Get(int64) (*Post, error)
	Delete(int64) error

	GetSlice(int64, int64) ([]*Post, error)
}

var (
	dup     sync.Mutex
	drivers = map[string]Store{}
)

func Register(name string, driver Store) {
	dup.Lock()
	defer dup.Unlock()
	if _, ok := drivers[name]; ok {
		panic("driver with name " + name + " already registered")
	}
	drivers[name] = driver
}

func Open(name string, setting string) (Store, error) {
	if driver, ok := drivers[name]; !ok {
		return nil, fmt.Errorf("Driver %s not found. Forgotten import?", name)
	} else {
		e := driver.Connect(setting)
		return driver, e
	}
}
