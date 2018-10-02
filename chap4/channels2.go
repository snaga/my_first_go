package main

import (
    "fmt"
    "time"
)

func main() {
    c := make(chan int)

    go func(s chan<- int) {
        fmt.Println("starts.")
        for i := 0; i < 3; i++ {
            time.Sleep(5 * time.Second)
            fmt.Println(i+1)
        }
        s <- 0
        fmt.Println("ends.")
    }(c)


    fmt.Println("waiting...")
    <-c
    fmt.Println("fin.")
}
