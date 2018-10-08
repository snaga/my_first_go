package main

import (
    "fmt"
    "strings"
)

func main() {
    s := "abcdefg"

    // substring match
    fmt.Println(strings.Contains(s, "cde"))
    fmt.Println(strings.Contains(s, "hij"))

    fmt.Println(strings.HasPrefix(s, "abc"))
    fmt.Println(strings.HasPrefix(s, "bcd"))
    fmt.Println(strings.HasSuffix(s, "def"))
    fmt.Println(strings.HasSuffix(s, "efg"))

    // upper/lower
    base := "aBcDeF"
    fmt.Println(strings.ToUpper(base))
    fmt.Println(strings.ToLower(base))

    // trim
    base2 := "!!!?!??   abcdef???!!!!"
    fmt.Println(base2)
    fmt.Println(strings.TrimLeft(base2, "!"))
    fmt.Println(strings.TrimRight(base2, "!?"))
    fmt.Println(strings.Trim(base2, "!? "))

    // replace
    base3 := "abcabcabcabc"
    fmt.Println(base3)
    fmt.Println(strings.Replace(base3, "bc", "yz", 2))
    fmt.Println(strings.Replace(base3, "abc", "xyz", -1))

    // split
    base4 := "ab::cd::efg::hij"
    fmt.Println(base4)
    fmt.Println(strings.Split(base4, "::"))
    fmt.Println(strings.SplitN(base4, "::", 3))

    // join
    a := []string{"C:", "work", "abc.txt"}
    fmt.Println(a)
    fmt.Println(strings.Join(a, "/"))
}
