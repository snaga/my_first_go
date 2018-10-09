package main

import (
    "encoding/xml"
    "fmt"
)

type Group struct {
    Name string         `xml:"name"`
    Companies []Company `xml:"company"`
}

type Company struct {
    Name string     `xml:"name"`
    Website Website `xml:"website"`
}

type Website struct {
    Name string `xml:",chardata"`
    URL string  `xml:"url,attr"`
}

func main() {
    head := Company{
        Name: "ABC株式会社",
        Website: Website{Name: "ABC公式ウェブサイト", URL: "http://abc.com"},
    }
    sol := Company{
        Name: "ABCソリューション株式会社",
        Website: Website{Name: "ソリューション事業について", URL: "http://abc.com/sol"},
    }

    data := new(Group)
    data.Name = "ABCグループ"
    data.Companies = []Company{head, sol}

    noIndent, err := xml.Marshal(data)
    if err != nil {
        fmt.Println("XML marshal error", err)
        return
    }
    fmt.Println(string(noIndent))

    withIndent, err := xml.MarshalIndent(data, "", "    ")
    if err != nil {
        fmt.Println("XML marshal error: ", err)
        return
    }
    fmt.Println(string(withIndent))
}

