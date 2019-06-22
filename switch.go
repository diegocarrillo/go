package main

import (
	"fmt"
	"time"
)

func main()  {
	weekdays()
}

func weekdays()  {
	
	Tiempo := time.Now() //We're getting the time var from here
	fmt.Println(Tiempo)
	switch CurrentDay := Tiempo.Weekday(); CurrentDay { //Using the var  Tiempo, we're assigning the current day to the var CurrentDay

	case 0:
		fmt.Println("sunday funday")
	case 6:
		fmt.Println("Why I don't have a beer at my hand, yet")
		fmt.Println(CurrentDay)
	default:
		fmt.Println("bah")
	}  

}
