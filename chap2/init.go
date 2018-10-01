package main

import "fmt"

var myval int = 17

func init() {
     fmt.Println(myval)
     fmt.Println("init")
}

func main() {
     fmt.Println("main")
}
