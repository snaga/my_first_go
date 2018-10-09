package main

import (
    "bytes"
    "encoding/csv"
    "fmt"
)

func main() {
    records := [][]string{
        []string{"りんご", "Apple", "バラ科"},
        []string{"みかん", "Orange", "ミカン科"},
        []string{"すいか", "Watermelon", "ウリ科"},
    }

    // Write()
    buf := new(bytes.Buffer)
    w := csv.NewWriter(buf)
    for _, record := range records {
        if err := w.Write(record); err != nil {
            fmt.Println("Write error: ", err)
            return
        }
        w.Flush()
    }
    fmt.Println(buf.String())

    // WriteAll()
    buf = new(bytes.Buffer)
    w = csv.NewWriter(buf)
    if err := w.WriteAll(records); err != nil {
        fmt.Println("Write error: ", err)
        return
    }
    fmt.Println(buf.String())
}
