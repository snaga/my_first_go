package main

import "fmt"

func main() {
    c := make(chan int)

    go func(s chan<- int) {
        fmt.Println("sending via channel.")
        for i := 0; i < 10; i++ {
            s <- i
        }
        close(s)
        fmt.Println("finish sending.")
    }(c)

    fmt.Println("receiving via channel.")
    for {
        val, ok := <-c
        if !ok {
            fmt.Println("channel closed.")
            break
        }
        fmt.Println(val)
    }
    fmt.Println("finish receiving.")
}
