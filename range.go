package main

import (
    "fmt"
)

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7}
    for i, num := range numbers {
        fmt.Printf("index:%d, num:%d \n", i, num)
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for i, c := range "gotonothing" {
        fmt.Println(i, c)
    }
}

