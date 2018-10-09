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
    xmlStr := `
<?xml version="1.0" encoding="UTF-8"?>
<group>
  <name>ABCグループ</name>
  <company>
    <name>ABC株式会社</name>
    <website url="http://abc.com">ABC公式ウェブサイト</website>
  </company>
  <company>
    <name>ABCソリューション株式会社</name>
    <website url="http://abc.com/sol">ソリューション事業について</website>
  </company>
</group>
`

    data := new(Group)
    if err := xml.Unmarshal([]byte(xmlStr), data); err != nil {
        fmt.Println("XML unmarshal error: ", err)
        return
    }

    fmt.Println(data.Name)
    fmt.Println(data.Companies[0].Name)
    fmt.Println(data.Companies[1].Website.Name)
    fmt.Println(data.Companies[1].Website.URL)
}

