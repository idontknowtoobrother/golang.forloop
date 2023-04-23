package main

import "fmt"

func main() {
	forLoopIndex()
	whileLoopPostPlusIndex()
	whileLoopPrePlusIndex()
}

func forLoopIndex() {
	for i := 0; i < 10; i++ {
		fmt.Printf("\nindex at %d", i)
	}
}

func whileLoopPostPlusIndex() {
	i := 0
	for i < 10 {
		fmt.Printf("\nindex at %d", i)
		i++ // post plus
	}
}

func whileLoopPrePlusIndex() {
	i := 0
	for i < 10 {
		i++ // pre plus
		fmt.Printf("\nindex at %d", i)
	}
}
