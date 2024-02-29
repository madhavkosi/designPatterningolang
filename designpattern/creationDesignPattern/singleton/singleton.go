package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}
var singleton *single

type single struct {
	val string
}

func NewSingleObject() *single {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()
		singleton = &single{val: "abc"}
		fmt.Printf("new created \n")
	} else {
		fmt.Printf("already created \n")
	}
	return singleton
}
