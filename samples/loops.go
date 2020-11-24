package main

import "fmt"

func main()  {
	pairs()
}

func pairs()  {
	
	number := 10

	for i := 0; i < number; i++{
		if i % 2 == 0 {
			fmt.Println("This is a pair number")
		} else {
			fmt.Println("This is a odd number")
		}
	}
}
