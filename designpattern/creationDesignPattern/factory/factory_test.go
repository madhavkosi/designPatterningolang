package factory

import (
	"fmt"

	"testing"
)

func TestFactory(t *testing.T) {
	a, err := PrinterFactory("color printer")
	if err != nil {
		fmt.Printf("following error %v\n", err)
		t.Errorf("failed")
		return
	}
	a.Print()
}
