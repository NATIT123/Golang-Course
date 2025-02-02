package main

import "fmt"

func ex2(str string) {
	frequent := make(map[rune]int)
	for _, c := range str {
		frequent[c]++
	}

	for c, count := range frequent {
		fmt.Printf("Character: %c, Frequency: %d\n", c, count)
	}

}
