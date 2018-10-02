package main

import (
    "fmt"
    "sync"
    "time"
)

func parallel(wg *sync.WaitGroup) {
    fmt.Println("A")
    time.Sleep(1 * time.Millisecond)
    fmt.Println("B")
    time.Sleep(1 * time.Millisecond)
    fmt.Println("C")

    wg.Done()
}

func main() {
    wg := new(sync.WaitGroup)

    for i := 0; i < 3; i++ {
        wg.Add(1)
        go parallel(wg)
    }

    wg.Wait()
}
