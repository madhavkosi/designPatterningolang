package main

import (
	"fmt"
)

// Function to compare two sets
func compareSets(set1 map[string][]string, set2 map[string]struct{}) {
	for value, data := range set1 {
		if _, found := set2[value]; !found {
			fmt.Println(data)
			fmt.Printf("Error: %s Should not be allowed to modify or delete as this is used in %s", value, data)
		}
	}
}

func main() {
	set1 := map[string][]string{
		"apple":  {"abc", ""},
		"banana": {},
		"cherry": {},
	}

	set2 := map[string]struct{}{
		"banana": {},
		"cherry": {},
		"date":   {},
	}

	compareSets(set1, set2)
}
