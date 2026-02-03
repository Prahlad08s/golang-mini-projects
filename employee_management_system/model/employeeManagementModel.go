package model

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Employee struct {
	Id     uint32  `validate:"required,min=1"`
	Name   string  `validate:"alphaspace"`
	Age    uint8   `validate:"required,min=18"`
	Salary float64 `validate:"required"`
}

type Department struct {
	Name         string      `validate:"alphaspace"`
	EmployeeList []*Employee `validate:"required"`
}

func (department *Department) ValidateDepartmentCredentials() error {
	validate := validator.New()
	err := validate.Struct(department)
	if err != nil {
		return err
	}
	return nil
}

func (employee *Employee) ValidateEmployeeCredentials() error {
	validate := validator.New()
	err := validate.Struct(employee)
	if err != nil {
		return err
	}
	return nil
}

func (department *Department) CalculateAverageSalary() float64 {
	employeeTotalCount := len(department.EmployeeList)

	if employeeTotalCount == 0 {
		return 0
	}

	totalSalary := float64(0)
	for _, employee := range department.EmployeeList {
		totalSalary += employee.Salary
	}

	return totalSalary / float64(employeeTotalCount)
}

func (department *Department) OnboardEmployee(scanner *bufio.Scanner) error {
	employee := &Employee{}

	// Get Employee ID
	fmt.Print("Enter the ID of the employee: ")
	scanner.Scan()
	id, err := strconv.ParseUint(strings.TrimSpace(scanner.Text()), 10, 32)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a valid ID.")
		return errors.New("invalid id")
	}
	if id < 1 {
		fmt.Println("ID must be at least 1.")
		return errors.New("id must be at least 1")
	}
	employee.Id = uint32(id)

	// Get Employee Name
	fmt.Print("Enter the name of the employee: ")
	scanner.Scan()
	employee.Name = strings.TrimSpace(scanner.Text())

	// Get Employee Age with bounds checking
	fmt.Print("Enter the age of the employee: ")
	scanner.Scan()
	age, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Println("Invalid age. Please enter a valid age.")
		return errors.New("invalid age")
	}
	if age < 0 || age > 255 {
		fmt.Println("Age must be between 0 and 255.")
		return errors.New("age out of range")
	}
	if age < 18 {
		fmt.Println("Age must be at least 18.")
		return errors.New("age too young")
	}
	employee.Age = uint8(age)

	// Get Employee Salary with validation
	fmt.Print("Enter the salary of the employee: ")
	scanner.Scan()
	salary, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
	if err != nil {
		fmt.Println("Invalid salary. Please enter a valid salary.")
		return errors.New("invalid salary")
	}
	if salary <= 0 {
		fmt.Println("Salary must be greater than 0.")
		return errors.New("salary must be positive")
	}
	employee.Salary = salary

	// Validate employee credentials
	inputValidationError := employee.ValidateEmployeeCredentials()
	if inputValidationError != nil {
		fmt.Println("Input validation error:", inputValidationError)
		return errors.New("input validation error")
	}

	department.EmployeeList = append(department.EmployeeList, employee)
	fmt.Println("successfully added the employee to the department")
	return nil
}

func (department *Department) OffboardEmployee(employeeName string) error {
	employeeName = strings.TrimSpace(employeeName)
	for indexOfEmployee, employee := range department.EmployeeList {
		if employee != nil && strings.TrimSpace(employee.Name) == employeeName {
			department.EmployeeList = append(department.EmployeeList[:indexOfEmployee], department.EmployeeList[indexOfEmployee+1:]...)
			fmt.Println("successfully offboarded the employee from the department")
			return nil
		}
	}
	return errors.New("employee not found")
}
