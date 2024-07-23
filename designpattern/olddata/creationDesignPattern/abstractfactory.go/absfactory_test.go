package abstractfactory

import (
	"fmt"

	"testing"
)

func TestFactory(t *testing.T) {
	a, err := AbsFactory("simple printer")
	if err != nil {
		fmt.Printf("following error %v\n", err)
		t.Errorf("failed")
		return
	}
	a.GetPrinter().Print()
	a.GetScanner().Scan()
}

// Abstract Factory is a creational design pattern that lets you produce families of related objects without specifying their concrete classe
