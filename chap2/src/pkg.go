package main

import (
  "somepkg"
  "fmt"
)

func main() {
     fmt.Println("Hello")

     fmt.Println(somepkg.SomeVar)
     somepkg.SomeFunc()
     fmt.Println(somepkg.SomeVar)
}
