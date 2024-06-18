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

// }
// The Factory Pattern defines an interface or abstract class for creating objects,

//The Factory Method pattern is a design pattern used to create objects.
//It allows a class to have a method that creates objects, but the specific type of object is determined by subclasses.
