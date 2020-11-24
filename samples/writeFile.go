package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {
	WriteToFile()
}

func WriteToFile()  {
	
	Text := []byte("This is a very important text")

	data := ioutil.WriteFile("NewFile", Text, 0755)
	ShowError(data)
	fmt.Println("Text was added")
}

func ShowError(e error)  {
	if e != nil {
		panic(e)
	}
}
