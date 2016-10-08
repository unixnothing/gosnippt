package main

import (
    "fmt"
    "runtime"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        runtime.Gosched()  // yield the current goroutine and switch to another.
        fmt.Println(s)
    }
}

func main() {
    go say("world")
    say("hello")
    //time.Sleep(time.Second)
}
