package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func returnGreetingFunc (salutation string) (greetingFunc func (string)) {
	greetingFunc = func (name string) {
		fmt.Printf("%v, %v", salutation, name)
	}
	return
}

func main() {
	newFunc := returnGreetingFunc("Good morning")
	newFunc("Skyler")
}

func intro () {
	scanner := bufio.NewScanner(os.Stdin)
	
	// Name as String
	fmt.Print("What is your name?: ")
	scanner.Scan()
	name := scanner.Text()
	
	fmt.Print("What is your age?: ")
	scanner.Scan()
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