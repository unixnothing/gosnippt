package main

import (
    "sort"
    "fmt"
)

func main() {
    strs := []string{"a", "z", "fdasf", "ok", "fdass"}
    sort.Strings(strs)
    fmt.Println(strs)

    ints := []int{12, 13, 3, 6, 2, 101}
    sort.Ints(ints)
    fmt.Println(ints)
}
