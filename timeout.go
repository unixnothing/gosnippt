package main

import (
    "time"
    "fmt"
)

// select timeout pattern
func main() {

    // eg 1 : timeout
    c1 := make(chan string)
    go func() {
        time.Sleep(time.Second * 2)
        c1 <- "f1 done"
    }()

    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(time.Second):
        fmt.Println("timeout 1")
    }


    // eg2: not timeout
    c2 := make(chan string)
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "f2 done"
    }()

    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(time.Second * 3):
        fmt.Println("timeout 2")

    }

}


