package main

import "fmt"

func plusOne(vals []int) {
     for i := 0; i < len(vals); i++ {
          vals[i] += 1
     }
}


func main() {
     alpha := [5]string{"a", "b", "c", "d", "e"}

     var slice1 []string

     slice1 = alpha[:]
     fmt.Println(slice1)

     slice2 := alpha[1:4]
     fmt.Println(slice2)

     slice3 := alpha[4:]
     fmt.Println(slice3)

     slice4 := alpha[:4]
     fmt.Println(slice4)

     num := [5]int{1, 2, 3, 4, 5}

     fmt.Println(num)
     plusOne(num[:])
     fmt.Println(num)

     s1 := num[1:4]
     fmt.Println("s1 = ", s1)
     fmt.Println("len = ", len(s1))
     fmt.Println("cap = ", cap(s1))

     s2 := s1[1:4]
     fmt.Println("s2 = ", s2)
     fmt.Println("len = ", len(s2))
     fmt.Println("cap = ", cap(s2))
}
