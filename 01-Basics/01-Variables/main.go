package main

import "fmt"

func main() {
	var firstName string = "Ahmed"
	var age = 26

	// Short Declaration
	jobTitle := "Backend Developer"
	salary := 5000.50
	isHired := true

	//Constants
	const CompanyName = "Tech Corp"

	// print data
	fmt.Println("=== Employee System ===")
	fmt.Println("Company:", CompanyName)
	fmt.Println("Name:", firstName)
	fmt.Println("Age:", age)
	fmt.Println("Job:", jobTitle)

	// Printf   (Formatted Print)
	// %f معناها رقم عشري، و .2 معناها علامتين بس بعد الصفر
	fmt.Printf("Salary: %.2f EGP\n", salary)
	fmt.Println("Hired Status:", isHired)
}
