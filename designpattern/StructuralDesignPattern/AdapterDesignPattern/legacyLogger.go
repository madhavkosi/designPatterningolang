package adapterdesignpattern

import "fmt"

type LegacyLogger struct {
}

func (l *LegacyLogger) LogMessage(message string) {
	fmt.Printf("message %s ", message)
}
