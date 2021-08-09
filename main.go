package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
)

func main () {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type a cool message for this Go code to read: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("You typed: %v", input)


	fmt.Println()
}