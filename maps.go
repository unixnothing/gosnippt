package main

import "fmt"

func main() {
    m := make(map[string]int)
    m["key1"] = 1
    m["key2"] = 1

    fmt.Println(m)
    fmt.Println(len(m))

    for k, v := range m {
        fmt.Println(k)
        fmt.Println(v)
    }

    delete(m, "key2")
    fmt.Println(m)

    _, prs := m["key1"]
    fmt.Println("key1 present:", prs)             // test whether the key is present.

    _, prs = m["key2"]
    fmt.Println("key2 present:", prs)

    n := map[int]string{1:"onnnne", 2:"twwwo"}    // declare and init.
    fmt.Println(n)
}
