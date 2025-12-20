package main

import "fmt"

func main() {
	array := [5]float64{30.0, 42.5, 28.0, 40.5, 35.5}
	var highTemps []float64

	for _, v := range array {
		if v >= 40.0 {
			highTemps = append(highTemps, v)
		}
	}

	fmt.Println("High temperature days:", highTemps)

}
