package main

import (
	"fmt"
)

func main() {
	forLoopIndex()
	whileLoopPostPlusIndex()
	whileLoopPrePlusIndex()

	players := []string{"John", "Doe", "Jane"}
	forEachWatchValue(players)
	forEachWatchOnlyValue(players)
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

func forEachWatchValue(data []string) {
	for index, value := range data {
		fmt.Printf("\nindex at %d, value is %s", index, value)
	}
}

func forEachWatchOnlyValue(data []string) {
	for _, value := range data {
		fmt.Printf("\nitem: %s", value)
	}
}
