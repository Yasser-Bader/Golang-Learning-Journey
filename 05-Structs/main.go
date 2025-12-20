package main

import "fmt"

type Employee struct {
	Name      string
	Age       int
	Salary    float64
	IsManager bool
}

func main() {
	emp := Employee{
		Name:      "Ali",
		Age:       24,
		Salary:    8000.5,
		IsManager: true,
	}
	fmt.Printf("Employee %v is %v years old and salary is %v \n ", emp.Name, emp.Age, emp.Salary)
	fmt.Println(emp)
	emp.Salary = 9000
	fmt.Println("Edited Salary ", emp.Salary)

}
