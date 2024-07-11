package main

import (
	"fmt"
)

// LogLevel represents a logging level.
type LogLevel int

// Define logging levels.
const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

// String converts the LogLevel to its string representation.
func (l LogLevel) String() string {
	return [...]string{"DEBUG", "INFO", "WARN", "ERROR"}[l]
}

// logMessage logs a message with a given log level.
func logMessage(level LogLevel, message string) {
	fmt.Printf("[%s] %s\n", level.String(), message)
}

func main() {
	// Create variables of type LogLevel.
	currentLevel := INFO
	debugLevel := DEBUG
	warnLevel := WARN
	errorLevel := ERROR

	// Log messages with different levels.
	logMessage(currentLevel, "This is an info message.")
	logMessage(debugLevel, "This is a debug message.")
	logMessage(warnLevel, "This is a warning message.")
	logMessage(errorLevel, "This is an error message.")
}
