package main

import (
	"bufio"
	"employee_management_system/business"
	"employee_management_system/model"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("---------------------------------------")
	fmt.Println("Welcome to EmployeeManagement Portal")
	fmt.Println("---------------------------------------")

	departmentList := make([]*model.Department, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("---------------------------------------")
		fmt.Println("1. Add a new department")
		fmt.Println("2. Onboard a new employee to a department")
		fmt.Println("3. Offboard an employee from a department")
		fmt.Println("4. Calculate the average salary of a department")
		fmt.Println("5. Show the department list")
		fmt.Println("6. Show the employee details of a department")
		fmt.Println("7. Exit")

		fmt.Print("Enter your choice: ")
		scanner.Scan()
		choiceStr := strings.TrimSpace(scanner.Text())
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid choice. Please enter a number between 1-7.")
			continue
		}

		switch choice {
		case 1:
			department := &model.Department{
				EmployeeList: make([]*model.Employee, 0),
			}
			fmt.Print("Enter the name of the department: ")
			scanner.Scan()
			department.Name = strings.TrimSpace(scanner.Text())
			if department.Name == "" {
				fmt.Println("Department name cannot be empty.")
				continue
			}
			inputValidationError := department.ValidateDepartmentCredentials()
			if inputValidationError != nil {
				fmt.Println("Input validation error:", inputValidationError)
				continue
			}
			departmentList = append(departmentList, department)
			fmt.Println("successfully added the department")
			continue

		case 2:
			fmt.Println("Employee to be added in which department?")
			scanner.Scan()
			departmentName := strings.TrimSpace(scanner.Text())
			business.OnboardEmployeeToDepartment(departmentList, departmentName, scanner)
			continue

		case 3:
			fmt.Println("Employee to be offboarded from which department?")
			scanner.Scan()
			departmentName := strings.TrimSpace(scanner.Text())
			business.OffboardEmployeeFromDepartment(departmentList, departmentName, scanner)
			continue

		case 4:
			fmt.Println("Department to calculate the average salary of?")
			scanner.Scan()
			departmentName := strings.TrimSpace(scanner.Text())
			averageSalary := business.CalculateAverageSalary(departmentList, departmentName)
			if averageSalary > 0 {
				fmt.Printf("The average salary of the department %s is %.2f\n", departmentName, averageSalary)
			}
			continue

		case 5:
			business.ShowDepartmentList(departmentList)
			continue

		case 6:
			fmt.Println("Department to show the employee details of?")
			scanner.Scan()
			departmentName := strings.TrimSpace(scanner.Text())
			business.ShowDepartmentEmployeeList(departmentList, departmentName)
			continue

		case 7:
			fmt.Println("Thank you for using EmployeeManagement Portal. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please enter a valid choice.")
			continue

		}
	}
}
