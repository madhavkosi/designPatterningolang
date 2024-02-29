package main

import "fmt"

// Printer defines the behavior of a printer.
type Printer interface {
	Print()
}
type SimplePrinter struct{}

// Print prints using the simple printer.
func (sp SimplePrinter) Print() {
	fmt.Println("Printing...")
}

type ColourfulplePrinter struct{}

// Print prints using the simple printer.
func (cp ColourfulplePrinter) Print() {
	fmt.Println("Printing Colourful...")
}
