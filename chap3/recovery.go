package main

import "fmt"

func func1(do_panic bool) {
     defer func() {
     	   fmt.Println("defer start.")
	   if err := recover(); err != nil {
	      fmt.Println("in recovery")
	   }
	   fmt.Println("defer end.")
     }()
     if do_panic {
     	  fmt.Println("panic")
          panic("panic caught.")
     }
}

func main() {
     func1(false)
     func1(true)
}
