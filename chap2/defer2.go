package main

import (
       "fmt"
       "os"
)

func main() {
     file, err := os.Open("defer2.go")
     if err != nil {
     	fmt.Println("File open error: ", err)
	return
     }
     defer file.Close()

     buf := make([]byte, 1024)
     for {
     	 n, err := file.Read(buf)
	 if err != nil {
	    fmt.Println("File read error: ", err)
	    return
	 }
	 if n == 0 {
	    break
	 }
	 fmt.Print(buf[:n])
     }
}
