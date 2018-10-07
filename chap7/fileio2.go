package main

import (
    "os"
    "fmt"
)

func main() {
    file, err := os.Create(`./fileio2.txt`)
    if err != nil {
        fmt.Println("os.Create() failed.")
        return
    }
    defer file.Close()

    output := "testmessage"
    file.Write(([]byte)(output))
}
