package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"vote-eligibility-checker/models"
)

func main() {

	employee := models.Person{}

	fmt.Println("Welcome to Vote Eligibility Checker")
	fmt.Println("--------------------------------")

	fmt.Print("Please enter your name:")
	inputScanner := bufio.NewScanner(os.Stdin)
	inputScanner.Scan()
	employee.Name = inputScanner.Text()

	fmt.Print("Please enter your age:")
	inputScanner.Scan()
	age, err := strconv.Atoi(inputScanner.Text())
	if err != nil {
		fmt.Println("Invalid age. Please enter a valid age.")
		return
	}
	employee.Age = uint8(age)

	inputValidationError := employee.ValidatePersonCredentials()
	if inputValidationError != nil {
		fmt.Println("Input validation error:", inputValidationError)
		return
	}

	for {

		fmt.Println("--------------------------------")
		fmt.Printf("Great! %s, let me know what you want to do next?\n", employee.Name)
		fmt.Println("1. Introduce myself")
		fmt.Println("2. Update my age")
		fmt.Println("3. Check my voting eligibility")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice:")
		inputScanner.Scan()
		userChoice, err := strconv.Atoi(inputScanner.Text())
		if err != nil {
			fmt.Println("Invalid choice. Please enter a valid number.")
			continue
		}

		switch userChoice {
		case 1:
			employee.IntroduceMe()
		case 2:
			fmt.Print("Please enter your new age:")
			fmt.Scanf("%d", &employee.Age)
			fmt.Println("Age is updated successfully.")
		case 3:
			employee.CheckVotingEligigbility()
		case 4:
			fmt.Println("Thank you for using Vote Eligibility Checker. Goodbye!")
			return
		}
	}

}
