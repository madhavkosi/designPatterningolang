package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const BufferSize = 5

type ProducerConsumer struct {
	Data       chan int
	ProducerWg sync.WaitGroup
	ConsumerWg sync.WaitGroup
}

func NewProducerConsumer() *ProducerConsumer {
	return &ProducerConsumer{
		Data: make(chan int, BufferSize),
	}
}

func (pc *ProducerConsumer) Producer(id int, items int) {
	defer pc.ProducerWg.Done()
	for i := 0; i < items; i++ {
		num := rand.Intn(100)
		fmt.Printf("Producer %d produced %d\n", id, num)
		pc.Data <- num
	}
}

func (pc *ProducerConsumer) Consumer(id int) {
	defer pc.ConsumerWg.Done()
	for num := range pc.Data {
		fmt.Printf("Consumer %d consumed %d\n", id, num)
	}
}

func (pc *ProducerConsumer) CloseChannel() {
	close(pc.Data)
}

func main() {
	pc := NewProducerConsumer()

	// Start Producers
	for i := 1; i <= 2; i++ {
		pc.ProducerWg.Add(1)
		go pc.Producer(i, 5)
	}

	// Start Consumers
	for i := 1; i <= 3; i++ {
		pc.ConsumerWg.Add(1)
		go pc.Consumer(i)
	}

	// Wait for all producers to finish
	pc.ProducerWg.Wait()

	// Once all producers are done, close the channel
	pc.CloseChannel()

	// Wait for all consumers to finish consuming
	pc.ConsumerWg.Wait()
}

package main

import (
	"fmt"
	"sync"
)

func printOdd(oddChan, evenChan chan struct{}, wg *sync.WaitGroup, idx int) {
	defer wg.Done()
	for i := 1; i <= idx; i += 2 {
		<-oddChan // Wait for the signal from the even number goroutine
		fmt.Printf("Odd: %d\n", i)
		if i != idx {
			evenChan <- struct{}{}
		}
	}
}

func printEven(oddChan, evenChan chan struct{}, wg *sync.WaitGroup, idx int) {
	defer wg.Done()
	for i := 2; i <= idx; i += 2 {
		<-evenChan // Wait for the signal from the odd number goroutine
		fmt.Printf("Even: %d\n", i)
		if i != idx {
			oddChan <- struct{}{}
		}
		// Signal the odd number goroutine to proceed
	}
}

func main() {
	oddChan := make(chan struct{})
	evenChan := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(2)
	go printOdd(oddChan, evenChan, &wg, 5)
	go printEven(oddChan, evenChan, &wg, 5)

	oddChan <- struct{}{}

	wg.Wait()
}
