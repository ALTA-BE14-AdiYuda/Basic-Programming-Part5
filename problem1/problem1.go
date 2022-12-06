package main

import "fmt"

func Compare(a, b string) string {
	// your code here
	var sama string
	if len(a) < len(b) {
		sama = a
	} else {
		sama = b
	}

	return sama
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}
