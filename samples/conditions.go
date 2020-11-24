package main

import "fmt"

func main()  {
	conditions()	
}

func conditions()  {
	
	grades := 69

	if grades < 70 {
		fmt.Println("you have no pass the test")
	} else if grades > 70 && grades <= 80{
		fmt.Println("You have pass the test")
	} else {
		fmt.Println("You are a NERF")
	}

}
