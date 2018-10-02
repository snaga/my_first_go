package main

import (
    "fmt"
    "sync"
    "time"
)

func parallel(wg *sync.WaitGroup, mt *sync.Mutex) {
    mt.Lock()
    defer mt.Unlock()

    fmt.Println("A")
    time.Sleep(1 * time.Millisecond)
    fmt.Println("B")
    time.Sleep(1 * time.Millisecond)
    fmt.Println("C")

    wg.Done()
}

func main() {
    wg := new(sync.WaitGroup)
    mt := new(sync.Mutex)

    for i := 0; i < 3; i++ {
        wg.Add(1)
        go parallel(wg, mt)
    }

    wg.Wait()
}
