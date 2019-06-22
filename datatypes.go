package main

import "fmt"

func main()  {
    var MyCoffe = Coffe{name:"Late", price:14.5, milk:2, sugar:true}

    fmt.Print(MyCoffe)
}

type Coffe struct {
    name string
    price float64
    milk int
    sugar bool
}
