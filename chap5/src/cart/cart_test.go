package cart

import (
//    "fmt"
    "testing"
)

/*
func TestMain(m *testing.M) {
    fmt.Println("hello")
}
*/

func TestAddAndGetProductsInCart(t *testing.T) {
    c := New()
    c.Add("りんご")
    c.Add("みかん")
//    c.Add("バナナ")

    products := c.GetAll()
    if len(products) != 2 {
        t.Fatalf("商品の数が違う %d", len(products))
    }
    if products[0] != "りんご" && products[1] != "りんご" {
        t.Error("りんご not found.")
        t.Log("カートの中身: ", products)
    }
    if products[0] != "みかん" && products[1] != "みかん" {
        t.Error("みかん not found.")
        t.Log("カートの中身: ", products)
    }
}
