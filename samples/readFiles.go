package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {
	
	readFile()

}

func readFile()  {
	
	Data, ErrorReading := ioutil.ReadFile("Poem.txt")
	
	showError(ErrorReading)

	fmt.Println(string(Data))

}

func showError(e error)  {
	
	if e != nil {
		panic(e)
	}

}
