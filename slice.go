package main

import "fmt"

func main() {
    s := make([]string, 10)    // make slice with init capacity
    fmt.Println(s)

    s[0] = "a"
    s[1] = "b"
    fmt.Println(s)
    fmt.Println("len:", len(s))

    s = append(s, "c", "d")          // capacity can be expand
    fmt.Println("len", len(s))
    fmt.Println(s)

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("c,the copyed slice: ", c)

    d := s[1:11]
    fmt.Println(d)
    d = s[:5]
    fmt.Println(d)
    d = s[5:]
    fmt.Println(d)

    t := []string{"g", "h", "i"}   // declare and init valuse
    fmt.Println("dcl:", t)
}
