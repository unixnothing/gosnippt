package main

import "fmt"

func main() {
    fmt.Println(fib(5))
}

func fib(n int) int {
    if n == 0 {
        return 1
    }

    return n * fib(n - 1)
}

