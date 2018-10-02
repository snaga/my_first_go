package main

import (
    "fmt"
    "strconv"
    "time"
)

func main() {
    fmt.Println("starts.")

    ch1 := make(chan int)
    ch2 := make(chan string)
    chend := make(chan struct{})

    go func(chint chan<- int, chstr chan<- string, end chan<- struct{}) {
        for i := 0; i < 10; i++ {
            if i % 2 == 0 {
                fmt.Println("to ch1")
                chint <- i
            } else {
                fmt.Println("to ch2")
                chstr <- "test " + strconv.Itoa(i)
            }
        }
        time.Sleep(1 * time.Second)
        close(end)
    }(ch1, ch2, chend)

    for {
        select {
        case val := <-ch1:
            fmt.Println("from ch1", val)
        case str := <-ch2:
            fmt.Println("from ch2", str)
        case <-chend:
            fmt.Println("fin.")
            return
        }
    }
}

