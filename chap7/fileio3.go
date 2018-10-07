package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    data, err := ioutil.ReadFile(`/etc/passwd`)
    if err != nil {
        fmt.Println("ioutil.ReadFile() failed.")
        return
    }
    fmt.Print(string(data))
}

