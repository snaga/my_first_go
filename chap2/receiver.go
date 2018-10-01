package main

import "fmt"

type SavingBox struct {
     money int
}

func NewBox() *SavingBox {
     return new(SavingBox)
}

func (s *SavingBox) Income(amount int) {
     s.money += amount
}

func (s *SavingBox) Break() int {
     lastMoney := s.money
     s.money = 0
     return lastMoney
}

func main() {
     box := NewBox()
     box.Income(100)
     box.Income(200)
     box.Income(500)

     fmt.Println(box.Break())
}
