package main

import "fmt"

func main()  {

    fmt.Print(greet())
}

func greet()(a string, b int) {
    a = "Hello "
    b = 75
    return
}
