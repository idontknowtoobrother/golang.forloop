package main

import "fmt"

func main() {
	forLoopIndex()
}

func forLoopIndex() {
	for i := 0; i < 10; i++ {
		fmt.Printf("\nindex at %d", i)
	}
}
