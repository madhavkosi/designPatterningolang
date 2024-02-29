package builder

import (
	"fmt"

	"testing"
)

func TestBuilder(t *testing.T) {
	c := NewCarBuilder().SetColor("Black").SetEngineType("electric").SetHasNavigation(false).SetHasSunRoof(true).Build()
	fmt.Printf("car specification %+v\n", c)
	d := NewCarBuilder().SetColor("Red").SetEngineType("Petrol").SetHasNavigation(true).SetHasSunRoof(true).Build()
	fmt.Printf("car specification %+v\n", d)
}
