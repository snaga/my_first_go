package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println(os.Args)

    if len(os.Args) != 4 {
        fmt.Println("Invalid command line parameters.")
        os.Exit(1)
    }

    fmt.Printf("Executable: %s\n", os.Args[0])
    fmt.Printf("Param 1: %s\n", os.Args[1])
    fmt.Printf("Param 2: %s\n", os.Args[2])
    fmt.Printf("Param 3: %s\n", os.Args[3])
}
