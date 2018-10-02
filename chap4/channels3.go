package main

import "fmt"

const goroutines = 5

func main() {
    counter := make(chan int)
    end := make(chan bool)

    // Goルーチンを5つ起動、共有のカウンタをそれぞれがインクリメントする。
    // 最後のGoルーチンはカウンタを書き戻さずに終了のシグナルを出す。
    for i := 0; i < goroutines; i++ {
        go func(counter chan int) {
            val := <-counter
            fmt.Println("counter = ", val, " -> ", val + 1)
            val++
            if val == goroutines {
                fmt.Println("ending")
                end <- true
            }
            counter <- val
        }(counter)
    }

    // 初期化して終了まで待機
    counter <- 0
    <-end

    fmt.Println("ends")
}
