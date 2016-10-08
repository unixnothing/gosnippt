package main

import (
    "time"
    "fmt"
)

func main() {
    timer5 := time.NewTimer(time.Second * 5)

    <-timer5.C //The <-timer5.C blocks on the timer’s channel C until it sends a value indicating that the timer expired.

    fmt.Println("timeOneSecond has expired")

    time10 := time.NewTimer(time.Second * 10)

    stop := time10.Stop()  // explicitly let it stop.
    if stop {
        fmt.Println("Time10 stopped")
    }
}
