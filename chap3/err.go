package main

import "fmt"

type MyError struct {
  msg string
}

func (e MyError) Error() string {
  return e.msg
}

func test_error() (val int, err error) {
  e := new(MyError)
  e.msg = "MyError caught"
  return -1, e
}

func main() {
     _, err := test_error()

     fmt.Println(err)
     fmt.Println(err.Error())
}
