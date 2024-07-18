package abstractfactory

import "fmt"

type Abstractfactory interface {
	GetPrinter() Printer
	GetScanner() Scanner
}

type SimpleScannerWithPrinter struct {
}

func (s SimpleScannerWithPrinter) GetPrinter() Printer {
	return BlackAndWhitePrinter{}
}
func (s SimpleScannerWithPrinter) GetScanner() Scanner {
	return SimpleScanner{}
}

func AbsFactory(factoryType string) (Abstractfactory, error) {
	switch factoryType {
	case "simple printer":
		return SimpleScannerWithPrinter{}, nil
	default:
		return nil, fmt.Errorf("factory type is not supported %s", factoryType)
	}

}
