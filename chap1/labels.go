package main

import "fmt"

func main() {
FOR_LABEL:
	for i := 0; i < 10; i++ {
	    switch {
	    case i == 3:
	    	 break FOR_LABEL
	    default:
	        fmt.Println(i)
	    }
	 }


LABEL1:
	for i := 0; i < 10; i++ {
	    fmt.Println("i = ", i)
	    for j := 0; j < 10; j++ {
	        fmt.Println("  j = ", j)
	    	if i == 0 && j == 5 {
		   fmt.Println("cnt1")
		   continue LABEL1
		} else if i == 1 && j == 1 {
		  fmt.Println("cnt2")
		  continue
		 }
	 }
  }


  for i := 0; i < 10; i++ {
      if i == 2 {
      	 goto LABEL
       }
      fmt.Println("i = ", i)
  }
LABEL:
}
