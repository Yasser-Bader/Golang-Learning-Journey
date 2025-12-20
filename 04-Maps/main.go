package main

import (
	"fmt"
)

func main() {
	employeeSalary := make(map[string]float64)
	employeeSalary["Ahmed"] = 6000.0
	employeeSalary["Ali"] = 5600.0
	employeeSalary["Yasser"] = 4500.0
	employeeSalary["Mona"] = 7200.0

	fmt.Printf("Salary Mona is  %v \n ", employeeSalary["Mona"])
	fmt.Println("Salaries", employeeSalary)

	delete(employeeSalary, "Ahmed")
	fmt.Print("Salaries ", employeeSalary)

}
