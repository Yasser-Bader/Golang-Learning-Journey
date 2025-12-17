package main
import "fmt"
func main()  {
	for i := 0; i <=20; i++ {
		if i%2==0{
			fmt.Printf("number %v is even \n",i)
			if i ==10{
				fmt.Println("📍 We reached the middle!")
			}
		}else{
			fmt.Printf("number %v is Odd \n",i)
		}
	}

}