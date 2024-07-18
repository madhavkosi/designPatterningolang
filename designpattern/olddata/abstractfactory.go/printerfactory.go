package abstractfactory

import "fmt"

type Printer interface {
	Print()
}

type BlackAndWhitePrinter struct {
}

func (s BlackAndWhitePrinter) Print() {
	fmt.Println("black and white printer")
}

type ColorPrinter struct {
}

func (s ColorPrinter) Print() {
	fmt.Println("color full  printer")
}
