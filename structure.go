package main

import "fmt"

func main() {
    fmt.Println(person{"Nothing", 28})

    p := person{name:"Chen", age:90}
    fmt.Println(p)

    pp := &p
    pp.age = 30
    fmt.Println(p)
}

type person struct {
    name string
    age  int
}
