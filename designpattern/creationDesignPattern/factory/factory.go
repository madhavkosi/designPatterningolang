package factory

import "fmt"

type BlackAndWhitePrinter struct {
}

func (s BlackAndWhitePrinter) Print() {
	fmt.Println("black printer")
}

type ColorPrinter struct {
}

func (s ColorPrinter) Print() {
	fmt.Println("color printer")
}

type Printer interface {
	Print()
}

func PrinterFactory(PrintType string) (Printer, error) {
	switch PrintType {
	case "black printer":
		return BlackAndWhitePrinter{}, nil
	case "color printer":
		return ColorPrinter{}, nil
	default:
		return nil, fmt.Errorf("Printer Type does not exist")
	}
}
