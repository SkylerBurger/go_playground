package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Example of a function closure
func returnGreetingFunc (salutation string) (greetingFunc func (string)) {
	greetingFunc = func (name string) {
		fmt.Printf("%v, %v", salutation, name)
	}
	return
}

func main() {
	//newFunc := returnGreetingFunc("Good morning")
	//newFunc("Skyler")

	// If a function is returned, you can immediately invoke it and pass args
	returnGreetingFunc("Good morning")("Skyler")
}

// Example of reading in user input from the CLI
func intro () {
	// Create a Scanner instance
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Print("What is your name?: ")
	// .Scan() when you want to read in user input
	scanner.Scan()
	// .Text() when you want to access the text that was scanned
	name := scanner.Text()
	
	fmt.Print("What is your age?: ")
	scanner.Scan()
	// Convert string type to integer type
	age, _ := strconv.ParseInt(scanner.Text(), 10, 8)
	fmt.Printf("Your name is %v and your age is %d years old.\n", name, age)
	
	year := 2021 - age
	fmt.Printf("Were you born in the year %v? (y/n): ", year)
	scanner.Scan()
	response := scanner.Text()

	if response == "y" {
		fmt.Println("My math skills are top notch!")
	} else {
		fmt.Println("Sorry I assumed you were so young. Your birthday must be coming up soon!")
	}

}