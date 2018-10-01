package main

import "fmt"

func func1() {
     defer func() {
     	   fmt.Println("defer 2")
     }()
     panic("Panic caught.")
}

func main() {
     defer func() {
     	   fmt.Println("defer 1")
     }()
     func1()
}
