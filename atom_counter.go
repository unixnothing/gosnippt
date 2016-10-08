package main

import (
    "fmt"
    "sync/atomic"
    "time"
)

func main() {
    num := uint64(0)
    for i := 0; i < 10000; i++ {
        go func() {
            atomic.AddUint64(&num, 1)
        }()
    }

    num2 := uint64(0)
    for i := 0; i < 10000; i++ {
        go func() {
            num2++
        }()
    }

    time.Sleep(time.Second)
    fmt.Println(num)
    fmt.Println(num2)
}
