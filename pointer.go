package main

import "fmt"

func main() {
     var pointer *int
     var n int = 100

     pointer = &n

     fmt.Println("nのアドレス: ", &n)
     fmt.Println("pointerの値: ", pointer)

     fmt.Println("nの値: ", n)
     fmt.Println("pointerの中身: ", *pointer)

     a, b := 10, 10
     called(a, &b)

     fmt.Println("値渡し: ", a)
     fmt.Println("ポインタ渡し: ", b)


     var p *int = new(int)
     *p = 13

     fmt.Println("p = ", p)
     fmt.Println("*p = ", *p)

     type myStruct struct {
     	  a int
	  b int
     }

     var my *myStruct = new(myStruct)

     (*my).a = 17
     (*my).b = 23

     fmt.Println("my = ", my)
     fmt.Println("*my = ", *my)
     fmt.Println("(*my).a = ", (*my).a)
     fmt.Println("(*my).b = ", (*my).b)
}

func called(a int, b *int) {
     a += 1
     *b += 1
}
