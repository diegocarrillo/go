package main

import "fmt"

func main()  {
	
	fruits()

}

func fruits()  {
	
	var nestedA[2][2]string 

	// This will be our first fruit and color
	nestedA[0][0] = "Apples"
	nestedA[0][1] = "red"

	//This will be our second fruit and it's color
	nestedA[1][0] = "grapes"
	nestedA[1][1] = "purple"

	fmt.Print(nestedA)
	fmt.Println(nestedA[1])
	fmt.Println(nestedA[0][1])

}
