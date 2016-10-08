package main

import "fmt"


// Use default in select to impl the non-blocking send,receives
func main() {
    message := make(chan string)

    select {
    case msg := <-message:
        fmt.Println("receive message:", msg)
    //default:
    //    fmt.Println("no message receive, prevent all goroutines dead and the message chan still have NO data")
    }

    fmt.Println("main finished")
}
