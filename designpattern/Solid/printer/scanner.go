package main

import "fmt"

// Scanner defines the behavior of a scanner.
type Scanner interface {
	Scan()
}

// SimpleScanner represents a basic scanner.
type SimpleScanner struct{}

// Scan scans using the simple scanner.
func (ss SimpleScanner) Scan() {
	fmt.Println("Scanning...")
}
