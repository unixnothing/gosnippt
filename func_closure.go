package main

import "fmt"

func main() {
    nextInt := intSeq();
    fmt.Println(nextInt)
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    //  the state(i) is unique to that particular function
    nextInt2 := intSeq()
    fmt.Println(nextInt2())
}

func intSeq() func() int {
    i := 0
    // return an anonymous function.
    return func() int {
        i += 1
        return i
    }
}


