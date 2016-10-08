package main

import (
    "time"
    "fmt"
)

func main() {
    ticker := time.NewTicker(time.Second * 1)
    go func() {
        for t := range ticker.C {
            fmt.Println("Ticker at", t)
        }
    }()

    time.Sleep(time.Second * 5)
}
