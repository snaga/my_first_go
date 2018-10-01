package main

import "fmt"

func main() {
     // append
     s1 := []int{1, 2, 3, 4, 5}
     fmt.Println("s1 = ", s1)

     s2 := append(s1, 6, 7)
     fmt.Println("s2 = ", s2)

     s3 := append(s1, s2...)
     fmt.Println("s3 = ", s3)

     // copy
     src1 := []int{1, 2}
     dest := []int{97, 98, 99}

     count := copy(dest, src1)
     fmt.Println("count = ", count)
     fmt.Println(dest)

     fmt.Println()

     src2 := []int{3}
     count = copy(dest[2:], src2)
     fmt.Println("count = ", count)
     fmt.Println(dest)

     // make
     ss1 := make([]string, 5, 10)
     fmt.Println("len=", len(ss1))
     fmt.Println("cap=", cap(ss1))

     ss2 := make([]string, 5)
     fmt.Println("len=", len(ss2))
     fmt.Println("cap=", cap(ss2))
}
