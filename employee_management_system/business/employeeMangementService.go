package business

import (
	"bufio"
	"employee_management_system/model"
	"fmt"
)

func OnboardEmployeeToDepartment(departmentList []*model.Department, departmentName string, scanner *bufio.Scanner) {
	for _, department := range departmentList {
		if department != nil && department.Name == departmentName {
			err := department.OnboardEmployee(scanner)
			if err != nil {
				fmt.Println("Error onboarding employee:", err)
			}
			return
		}
	}
	fmt.Println("Department not found")
}

func OffboardEmployeeFromDepartment(departmentList []*model.Department, departmentName string, scanner *bufio.Scanner) {
	for _, department := range departmentList {
		if department.Name == departmentName {
			fmt.Print("Enter the name of the employee to be offboarded: ")
			scanner.Scan()
			employeeName := scanner.Text()
			err := department.OffboardEmployee(employeeName)
			if err != nil {
				fmt.Println("Error offboarding employee:", err)
			} else {
				fmt.Println("Employee offboarded successfully")
			}
			return
		}
	}
	fmt.Println("Department not found")
}

func ShowDepartmentList(departmentList []*model.Department) {
	if len(departmentList) == 0 {
		fmt.Println("No departments found.")
		fmt.Println("--------------------------------")
		return
	}
	for _, department := range departmentList {
		if department != nil {
			fmt.Println("Department:", department.Name)
		}
	}
	fmt.Println("--------------------------------")
}

func ShowDepartmentEmployeeList(departmentList []*model.Department, departmentName string) {
	for _, department := range departmentList {
		if department != nil && department.Name == departmentName {
			if len(department.EmployeeList) == 0 {
				fmt.Println("No employees in this department.")
				return
			}
			fmt.Println("Employee List:")
			for i, employee := range department.EmployeeList {
				if employee != nil {
					fmt.Printf("  %d. ID: %d, Name: %s, Age: %d, Salary: %.2f\n",
						i+1, employee.Id, employee.Name, employee.Age, employee.Salary)
				}
			}
			return
		}
	}
	fmt.Println("Department not found")
}

func CalculateAverageSalary(departmentList []*model.Department, departmentName string) float64 {
	for _, department := range departmentList {
		if department.Name == departmentName {
			return department.CalculateAverageSalary()
		}
	}
	fmt.Println("Department not found")
	return 0
}
