package main

import (
       "fmt"
       "time"
)

func main() {
    hour := time.Now().Hour()
    if hour >= 6 && hour < 12 {
     	fmt.Println("朝です。")
    } else if hour < 19 {
      	fmt.Println("昼です。")
    } else {
      	fmt.Println("夜です。")
    }

    if hour := time.Now().Hour(); hour >= 6 && hour < 12 {
        fmt.Println("朝です。")
    } else if hour < 19 {
        fmt.Println("昼です。")
    } else {
        fmt.Println("夜です。")
    }

    dayOfWeek := "日"
    switch dayOfWeek {
    case "土":
    	 fmt.Println("大概は休みです。")
    case "日":
         fmt.Println("ほぼ間違いなく休みです。")
    default:
        fmt.Println("仕事です・・・。")
    }

    dayOfWeek2 := "土"
    switch dayOfWeek2 {
    case "土":
    	 fallthrough
    case "日":
         fmt.Println("休みです。")
    default:
        fmt.Println("仕事です・・・。")
    }

    dayOfWeek3 := "土"
    switch dayOfWeek3 {
    case "土", "日":
         fmt.Println("休みです。")
    default:
        fmt.Println("仕事です・・・。")
    }

    hour2 := time.Now().Hour()
    switch {
    case hour2 >= 6 && hour2 < 12:
    	 fmt.Println("朝です。")
    case hour2 < 19:
    	 fmt.Println("昼です。")
    default:
	 fmt.Println("夜です。")
    }
}
