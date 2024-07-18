package adapterdesignpattern

import "testing"

func TestAdapterDesignPatern(t *testing.T) {
	l := LegacyLogger{}
	la := LoggerAdapter{Adapter: &l}
	la.Log("hello how are you")
	nl := NewLogger{}
	nl.Log("hello how are you")
}
