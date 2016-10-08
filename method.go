package main

import "fmt"

func main() {
    r1 := rect{4, 6}
    fmt.Println("method area:", r1.area()) // Go automatically handles conversion between values and pointers for method calls
    fmt.Println("func area:", area(r1))

    fmt.Println("perim:", r1.perimeter())

    r2 := &r1
    fmt.Println(r2.perimeter()) // Go automatically handles conversion between values and pointers for method calls
    fmt.Println(r2.area())
}

type  rect struct {
    width, height int
}

func (r rect) perimeter() int {
    return (r.height + r.width) * 2
}

func (r *rect) area() int {
    return r.height * r.width
}

func area(r rect) int {
    return r.height * r.width
}
