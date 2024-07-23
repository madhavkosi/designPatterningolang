package abstractfactory

import "fmt"

type LaserScanner struct {
}

func (s LaserScanner) Scan() {
	fmt.Println("Laser Scanner")
}

type HandHeldScanner struct {
}

func (s HandHeldScanner) Scan() {
	fmt.Println("handHeld Scanner")
}

type SimpleScanner struct {
}

func (s SimpleScanner) Scan() {
	fmt.Println("Laser Scanner")
}

type Scanner interface {
	Scan()
}
