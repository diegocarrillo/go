package main

import "fmt"

func main()  {
	fruits()
}

func fruits()  {
	
	var MyFruits = []string{"grapes", "apples", "bananas"}

	MyFruits = append(MyFruits, "watermelons")

	fmt.Println(MyFruits)
	fmt.Println(MyFruits[3])
	fmt.Println(len(MyFruits))

}
