package main

import "fmt"

func main() {
     currencies := make(map[string]string)

     currencies["日本"] = "JPY"
     currencies["USA"] = "USD"
     currencies["EU"] = "EUR"
     currencies["中国"] = "CNY"

     fmt.Println(currencies["日本"])
     fmt.Println(currencies["中国"])

     fmt.Println("すべて取得")

     for country, currency := range currencies {
     	 fmt.Println(country, currency)
     }

     _, exist := currencies["ロシア"]
     if exist {
     	fmt.Println("ロシア: found")
     } else {
     	fmt.Println("ロシア: not found")
     }

     _, exist = currencies["日本"]
     if exist {
     	fmt.Println("日本: found")
     } else {
     	fmt.Println("日本: not found")
     }

     delete(currencies, "日本")

     _, exist = currencies["日本"]
     if exist {
     	fmt.Println("日本: found")
     } else {
     	fmt.Println("日本: not found")
     }

     currencies2 := map[string]string{
     		 "日本": "JPY",
		 "米国": "USD",
		 "EU": "EUR",
		 "中国": "CNY",
     }
     fmt.Println(currencies2)
}
