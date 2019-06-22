//We will define some functions here
package main

import "fmt"

func main()  {

    fmt.Print(mult(), multiplication(20, 20))


}

func mult()  int {
    return 100
}

func multiplication(num1, num2 int) int {

    return num1 * num2

}
