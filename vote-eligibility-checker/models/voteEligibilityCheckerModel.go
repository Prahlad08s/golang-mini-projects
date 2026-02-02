package models

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Person struct {
	Name string `validate:"required,min=2,max=100"`
	Age  uint8  `valiate:"required, gte=0"`
}

// Note: Methods can only be defined in the same package as the type they're attached to.
// This is the reason we are not defining these methods in business package.

func (person *Person) ValidatePersonCredentials() error {
	validate := validator.New()
	err := validate.Struct(person)
	if err != nil {
		return err
	}
	return nil
}

func (person *Person) IntroduceMe() {
	fmt.Printf("Hello there. My name is %s and I am %d years old.\n", person.Name, person.Age)
}

func (person *Person) CheckVotingEligigbility() {
	if person.Age >= 18 {
		fmt.Println("You are eligible to vote.")
	} else {
		fmt.Println("You are not eligible to vote.")
	}
}
