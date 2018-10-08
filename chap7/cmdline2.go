package main

import (
    "fmt"
    "flag"
)

var (
    intOpt  = flag.Int("i", 1234, "help for i option")
    boolOpt = flag.Bool("b", false, "help for b option")
    strOpt  = flag.String("s", "default", "help for s option")
)

func main() {
    flag.Parse()

    fmt.Println(*intOpt)
    fmt.Println(*boolOpt)
    fmt.Println(*strOpt)
}
