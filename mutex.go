package main

import (
    "sync"
    "fmt"
    "time"
)

func main() {
    mutex := &sync.Mutex{}

    num := uint64(0)
    for i := 0; i < 10000; i++ {
        go func() {
            mutex.Lock()
            num++
            mutex.Unlock()
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
