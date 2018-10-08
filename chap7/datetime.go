package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println(time.Now())

    // localtime
    t := time.Date(2015, 9, 13, 12, 35, 42, 123456789, time.Local)
    fmt.Println(t)

    // defining location
    loc := time.FixedZone("EST", -5*60*60)
    t = time.Date(2015, 9, 13, 12, 35, 42, 123456789, loc)
    fmt.Println(t)

    // pre-defined location
    loc2, err := time.LoadLocation("America/New_York")
    if err != nil {
        fmt.Println("time.LoadLocation() failed.")
        return
    }
    t = time.Date(2015, 9, 13, 12, 35, 42, 123456789, loc2)
    fmt.Println(t)

    // date/time format
    now := time.Now()
    fmt.Printf("Not formatted: %s\n", now.String())
    fmt.Printf("Formatted: %s\n", now.Format("2006/01/02 Mon 15:04:05"))

    // extracting date/time
    fmt.Printf("string: %s\n", now.String())
    fmt.Printf("year: %d\n", now.Year())
    fmt.Printf("month: %d\n", now.Month())
    fmt.Printf("day: %d\n", now.Day())
    fmt.Printf("weekday: %s\n", now.Weekday().String())
    fmt.Printf("hour: %d\n", now.Hour())
    fmt.Printf("minute: %d\n", now.Minute())
    fmt.Printf("second: %d\n", now.Second())
    fmt.Printf("nano-second: %d\n", now.Nanosecond())

    // date/time comparison
    base := time.Date(2015, 11, 1, 0, 0, 0, 0, time.Local)
    same := time.Date(2015, 11, 1, 0, 0, 0, 0, time.Local)
    before := time.Date(2015, 10, 31, 23, 59, 59, 0, time.Local)
    after := time.Date(2015, 11, 1, 0, 0, 1, 0, time.Local)

    fmt.Println(base.Equal(same))
    fmt.Println(base.Before(after))
    fmt.Println(base.After(before))

    // date/time calculation
    fmt.Println(base)
    fmt.Println(base.Add(7 * time.Hour))
    fmt.Println(base.Add(30 * time.Minute))
    fmt.Println(base.Add(-5 * time.Second))

    // subtract
    fmt.Println(base.Sub(before))

    duration := time.Duration(1)
    fmt.Println(base)
    fmt.Println(base.Add(duration))
}
