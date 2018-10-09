package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "strings"
)

func main() {
    lines := []string {
        "りんご,Apple,バラ科",
        "みかん,Orange,ミカン科",
        "すいか,Watermelon,ウリ科",
    }
    csvStr := strings.Join(lines, "\n")

    // Read()
    r := csv.NewReader(strings.NewReader(csvStr))
    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println("Read error: ", err)
            break
        }

        fmt.Printf("ja[%s] en[%s] category[%s]\n", record[0], record[1], record[2])
    }

    // ReadAll()
    r = csv.NewReader(strings.NewReader(csvStr))
    records, err := r.ReadAll()
    if err != nil {
        fmt.Println("Read error: ", err)
        return
    }
    for _, record := range records {
        fmt.Printf("ja[%s] en[%s] category[%s]\n", record[0], record[1], record[2])
    }
}
