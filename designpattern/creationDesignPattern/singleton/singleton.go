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

func (s single) values() {
	fmt.Println("Abc")
}

func NewSingleObject() *single {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleton == nil {
			singleton = &single{val: "abc"}
			fmt.Printf("new created \n")
		} else {
			fmt.Printf("already created \n")
		}
	} else {
		fmt.Printf("already created \n")
	}
	return singleton
}
