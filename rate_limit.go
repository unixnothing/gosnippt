package main

import (
    "time"
    "fmt"
)

func main() {
    // assume we have got 10 request to handle with
    requests := make(chan int, 10)
    for i := 0; i < 10; i++ {
        requests <- i
    }
    close(requests)

    limiter := time.Tick(time.Millisecond * 200)

    // handle every request
    for req := range requests {
        <-limiter  //By blocking on a receive from the limiter channel before serving each request, we limit ourselves to 1 request every 200 milliseconds.
        fmt.Println("handle request", req, time.Now())
    }

    // Only do limit control when it have been burst(reach 10 request).
    burstyLimiter := make(chan time.Time, 10) // use time to control the rate limit, every time come into the chan, it will handle a req.
    for i := 0; i < 10; i++ {
        burstyLimiter <- time.Now()
    }

    //every 200ms, will send a time into the chann, tells that the req can be handle.
    for t := range time.Tick(time.Millisecond * 200) {
        fmt.Println(burstyLimiter)
        burstyLimiter <- t
    }

    requests2 := make(chan int, 20)
    for i := 0; i < 20; i++ {
        requests2 <- i
    }
    close(requests2)

    for req := range requests2 {
        <-burstyLimiter
        fmt.Println("handle request2", req, time.Now())
    }
}
