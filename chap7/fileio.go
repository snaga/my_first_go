package main

import (
    "fmt"
    "os"
)

const BUFSIZE = 1024

func main() {
    file, err := os.Open(`/etc/passwd`)
    if err != nil {
        fmt.Println("os.open() failed.")
        return
    }
    defer file.Close()

    buf := make([]byte, BUFSIZE)
    for {
        n, err := file.Read(buf)
        if n == 0 {
            break
        }
        if err != nil {
            fmt.Println("file.Read() failed.")
            return
        }

        fmt.Print(string(buf[:n]))
    }
}
