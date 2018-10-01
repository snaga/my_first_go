package main

import "fmt"

type Animal interface {
     Cry()
}

type Dog struct {}
func (d *Dog) Cry() {
     fmt.Println("わん")
}

type Cat struct {}
func (c *Cat) Cry() {
     fmt.Println("にゃー")
}

func make_cry(a Animal) {
     a.Cry()
}

func make_cry_2(obj interface{}) {
     a, ok := obj.(Animal)
     if !ok {
     	fmt.Println("Not an Animal object")
     	return
     }
     a.Cry()
}

func main() {
     dog := new(Dog)
     cat := new(Cat)
     make_cry(dog)
     make_cry(cat)

     make_cry_2(dog)
     make_cry_2(cat)

     foo := "foo"
     make_cry_2(foo)
}
