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

	fmt.Println("\nfound:", findTillFoundString(players, "Doe"))
	fmt.Println("\nfound:", findTillFoundString(players, "Keng"))

	fmt.Println("\nfound:", isStrInData(players, "Doe"))
	fmt.Println("\nfound:", isStrInData(players, "Keng"))

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

func findTillFoundString(data []string, findStr string) (found bool) {
	found = false
	for _, str := range data {
		found = str == findStr
		if found {
			break
		}
	}
	return found
}

func isStrInData(data []string, findStr string) (found bool) {
	found = false
	length := len(data) - 1
	for !found && length >= 0 {
		found = data[length] == findStr
		length--
	}
	return found
}
