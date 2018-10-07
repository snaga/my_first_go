package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open(`/etc/passwd`)
    if err != nil {
        fmt.Println("os.Open() failed.")
        return
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    for i := 1; sc.Scan(); i++ {
        if err := sc.Err(); err != nil {
            fmt.Println("bufio.NewScanner.Scan() failed.")
            break
        }
        fmt.Printf("LINE %d: %s\n", i, sc.Text())
    }
}
