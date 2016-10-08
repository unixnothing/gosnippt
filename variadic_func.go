package main

import (
    "fmt"
    "reflect"
)

func main() {
    a := sum(1, 2)
    b := sum(1, 2, 3)
    c := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
    fmt.Println(a, b, c)

    nums:=[]int{1,2,3,4,5}
    d:=sum(nums...)  // attention.
    fmt.Println(d)
}

// var args function
func sum(num ...int) int {
    fmt.Println(reflect.TypeOf(num))
    sum := 0
    for _, single := range num {
        sum += single
    }
    return sum
}


