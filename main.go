package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create Input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type something: ")
	scanner.Scan()

	input := scanner.Text()

	// create character container
	var characterContainer = map[string]int{}

	for i := range input {
		var a string = string(byte(input[i]))
		characterContainer[a] = 0
		fmt.Printf("%v \n", a)
	}

	// find and count each character
	for i := range input {
		var a string = string(byte(input[i]))
		findCharacter(characterContainer, a)
	}

	fmt.Println("====================================================")
	fmt.Println(characterContainer)
}

func findCharacter(characters map[string]int, a string) {
	for key := range characters {
		if key == a {
			characters[key] += 1
		}
	}
}
