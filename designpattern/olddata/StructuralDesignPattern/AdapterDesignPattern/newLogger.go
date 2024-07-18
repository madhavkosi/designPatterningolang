package adapterdesignpattern

import "fmt"

type NewLogger struct {
}

func (l *NewLogger) Log(message string) {
	fmt.Printf("message %s", message)
}
