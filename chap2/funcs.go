package main

import "fmt"

func DisplayHello() {
     fmt.Println("Hello!")
}

func DisplaySum(left int, right int) {
     fmt.Println(left + right)
}

func DisplaySumAll(values ...int) {
     sum := 0
     for _, value := range values {
     	 sum += value
     }
     fmt.Println(sum)
}

func Sum(left int, right int) int {
     return left + right
}

func Div(left int, right int) (int, int) {
     return left / right, left % right
}


func main() {
     DisplayHello()

     DisplaySum(2, 5)

     DisplaySumAll(2, 5, 8, 11)

     fmt.Println(Sum(2, 5))

     fmt.Println(Div(19, 4))
}
