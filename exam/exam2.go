package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Employee struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Salary     int    `json:"salary"`
	Experience int    `json:"experience"`
	Age        int    `json:"age"`
}

func main() {
	data, _ := os.ReadFile("employees.json")

	var employees []Employee
	json.Unmarshal(data, &employees)

	sort.Slice(employees, func(i, j int) bool {
		return employees[i].Age < employees[j].Age
	})

	orderByAgeFile, _ := os.Create("orderByAge.txt")
	defer orderByAgeFile.Close()

	for _, employee := range employees {
		orderByAgeFile.WriteString(fmt.Sprintf("\"first_name\": \"%s\" \"second_name\": \"%s\" \"salary\": %d \"experience\": %d \"age\": %d\n",
			employee.FirstName, employee.SecondName, employee.Salary, employee.Experience, employee.Age))
	}

	sort.Slice(employees, func(i, j int) bool {
		return employees[i].Salary > employees[j].Salary
	})

	topSalaryEmployeesFile, _ := os.Create("topSalaryEmployees.txt")
	defer topSalaryEmployeesFile.Close()

	for i := 0; i < 3 && i < len(employees); i++ {
		topSalaryEmployeesFile.WriteString(fmt.Sprintf("\"first_name\": \"%s\" \"second_name\": \"%s\" \"salary\": %d \"experience\": %d \"age\": %d\n",
			employees[i].FirstName, employees[i].SecondName, employees[i].Salary, employees[i].Experience, employees[i].Age))
	}
}
