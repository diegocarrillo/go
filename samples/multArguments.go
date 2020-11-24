package main

import "fmt"

func main()  {
    class("Diego", "Hugo", "Rogelio")
}

func class(clasmates ...string) {

    for _, val := range clasmates {
    fmt.Println(val)
    }

}
