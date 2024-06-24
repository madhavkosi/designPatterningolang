package main

import "fmt"

// PhotoCopier represents a photocopier that can print and scan.
type PhotoCopier struct {
	Printer
	Scanner
}
type BasicPrinter struct {
	Printer
}

func main() {
	// SimplePrinter and SimpleScanner are independent and can be used separately.
	fmt.Println("Simple Printer:")
	simplePrinter := SimplePrinter{}
	simplePrinter.Print()

	fmt.Println("\nSimple Scanner:")
	simpleScanner := SimpleScanner{}
	simpleScanner.Scan()

	// PhotoCopier is a MultiFunctionDevice that implements both Printer and Scanner.
	fmt.Println("\nPhotocopier:")
	photoCopier := PhotoCopier{
		Printer: SimplePrinter{},
		Scanner: SimpleScanner{},
	}
	photoCopier.Print()
	photoCopier.Scan()
	BP := BasicPrinter{Printer: simplePrinter}
	BP.Print()
	s := BasicPrinter{Printer: ColourfulplePrinter{}}
	s.Print()
}
