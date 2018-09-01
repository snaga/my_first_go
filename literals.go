package main

import "fmt"

type Score int

func main() {
     // integers
     a := 1234
     fmt.Println(a)
     b := 073
     fmt.Println(b)
     c := 0xa3
     fmt.Println(c)

     // floating points
     d := 4.1415
     fmt.Println(d)
     e := .25
     fmt.Println(e)
     f := 12.
     fmt.Println(f)
     g := 1.25e-3
     fmt.Println(g)

     // complex number
     h := 2i
     fmt.Println(h)

     // runes
     i := 'a'
     fmt.Println(i)
     j := 'あ'
     fmt.Println(j)
     k := '\n'
     fmt.Println(k)
     l := '\u12AB'
     fmt.Println(l)

     // raw
     fmt.Println(`abc`)
     fmt.Println(`\n`)
     fmt.Println(`ab
cd`)

     // interpreted
     fmt.Println("abc")
     fmt.Println("ab\ncd")
     fmt.Println("\u3042\u3044\u3046")
}
