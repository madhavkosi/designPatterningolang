package main

import "fmt"

// SupportHandler is the handler interface
type SupportHandler interface {
	SetNext(handler SupportHandler)
	HandleRequest(level int, message string)
}

// BaseHandler is the base struct that implements the common logic for setting the next handler
type BaseHandler struct {
	next SupportHandler
}

func (h *BaseHandler) SetNext(next SupportHandler) {
	h.next = next
}

func (h *BaseHandler) HandleRequest(level int, message string) {
	if h.next != nil {
		h.next.HandleRequest(level, message)
	}
}

// LevelOneSupport is a concrete handler for Level 1 support
type LevelOneSupport struct {
	BaseHandler
}

func (h *LevelOneSupport) HandleRequest(level int, message string) {
	if level == 1 {
		fmt.Println("Level 1 Support: Handling request -", message)
	} else {
		h.BaseHandler.HandleRequest(level, message)
	}
}

// LevelTwoSupport is a concrete handler for Level 2 support
type LevelTwoSupport struct {
	BaseHandler
}

func (h *LevelTwoSupport) HandleRequest(level int, message string) {
	if level == 2 {
		fmt.Println("Level 2 Support: Handling request -", message)
	} else {
		h.BaseHandler.HandleRequest(level, message)
	}
}

// LevelThreeSupport is a concrete handler for Level 3 support
type LevelThreeSupport struct {
	BaseHandler
}

func (h *LevelThreeSupport) HandleRequest(level int, message string) {
	if level == 3 {
		fmt.Println("Level 3 Support: Handling request -", message)
	} else {
		h.BaseHandler.HandleRequest(level, message)
	}
}

func main() {
	levelOne := &LevelOneSupport{}
	levelTwo := &LevelTwoSupport{}
	levelThree := &LevelThreeSupport{}

	// Set up the chain: Level 1 -> Level 2 -> Level 3
	levelOne.SetNext(levelTwo)
	levelTwo.SetNext(levelThree)

	// Test the chain with different request levels
	levelOne.HandleRequest(1, "Password reset request")
	levelOne.HandleRequest(2, "Software installation request")
	levelOne.HandleRequest(3, "System outage report")
	levelOne.HandleRequest(4, "Unrecognized support request")
}
