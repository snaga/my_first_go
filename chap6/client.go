package main

import (
    "fmt"
    "net"
)

func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:1234")
    if err != nil {
        fmt.Printf("Dial error: %s\n", err)
        return
    }
    defer conn.Close()

    sendMsg := "Test Message.\n"
    conn.Write([]byte(sendMsg))
}
