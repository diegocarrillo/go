package main

import "fmt"

func main()  {
    fmt.Print(cows(1))
}

func cows(number int) (int, string){

  cowCounter := func() int {

  return number * 10
}
  return cowCounter(), " cows"

}
