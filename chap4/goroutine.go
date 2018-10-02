package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("main start.")

    fmt.Println("Calling as a regular function.")
    serialno()

    fmt.Println("Calling as a go routing.")
    go serialno()

    time.Sleep(1 * time.Second)
    fmt.Println("main end.")
}

func serialno() {
    fmt.Println("serialno starts.")
    for i := 0; i < 5; i++ {
        fmt.Println(i)
        time.Sleep(1 * time.Second)
    }
    fmt.Println("serialno ends.")
}
