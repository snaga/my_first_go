package main

import "fmt"

func main() {
     var month [12]string

     month[0] = "January"
     month[1] = "February"
     month[2] = "March"
     month[3] = "April"
     month[4] = "May"
     month[5] = "June"
     month[6] = "July"
     month[7] = "August"
     month[8] = "September"
     month[9] = "October"
     month[10] = "November"
     month[11] = "December"

     for i := 0; i < len(month); i++ {
     	 fmt.Printf("%d = %s\n", i+1, month[i])
     }

     month2 := [...]string{
          "January",
	  "February",
	  "March",
	  "April",
	  "May",
	  "June",
	  "July",
	  "August",
	  "September",
	  "October",
	  "November",
	  "December"}

     for i, m := range month2 {
     	 fmt.Printf("%d = %s\n", i+1, m)
     }
}
