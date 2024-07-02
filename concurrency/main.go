package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var zeroch chan bool
var oddch chan bool
var evench chan bool

func zero(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(0)

		if i%2 == 0 {
			oddch <- true
		} else {
			evench <- true
		}
		<-zeroch
		defer wg.Done()
	}
}
func odd(n int) {
	for i := 1; i <= n; i += 2 {
		<-oddch
		fmt.Println(i)
		defer wg.Done()
		zeroch <- true
	}
}

func even(n int) {
	for i := 2; i <= n; i += 2 {
		<-evench
		fmt.Println(i)
		defer wg.Done()
		zeroch <- true
	}
}

func main() {
	n := 2
	wg.Add(n * 2)
	go zero(n)
	go odd(n)
	go even(n)
	wg.Wait()
}
