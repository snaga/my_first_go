package main

import "fmt"

type MyVector struct {
    X int
    Y int
    z int
}

func main() {
    var v MyVector
    v.X = 2  // OK
    v.Y = 5  // OK
    v.z = 8  // OK

    fmt.Println(v)
}
