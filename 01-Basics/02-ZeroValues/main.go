package main

import "fmt"

func main() {
	// بنعرف متغيرات من غير ما نديها أي قيم ابتدائية
	var age int
	var salary float64
	var name string
	var isDeveloper bool

	fmt.Println("=== Go Zero Values ===")

	// هنطبع القيم عشان نشوف Go حطت فيهم إيه
	// %q مع الاسترينج بيطبعه بين علامتين تنصيص عشان يبان لو هو فاضي
	fmt.Printf("Age (int): %d\n", age)          // المتوقع: 0
	fmt.Printf("Salary (float64): %.2f\n", salary) // المتوقع: 0.00
	fmt.Printf("Name (string): %q\n", name)       // المتوقع: "" (نص فاضي)
	fmt.Printf("IsDeveloper (bool): %v\n", isDeveloper) // المتوقع: false

	// معلومة إضافية: معرفة نوع المتغير
	fmt.Printf("\nType of age is: %T\n", age)
}
