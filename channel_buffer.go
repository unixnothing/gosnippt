package main

import "fmt"

func main() {
    messages := make(chan string, 2)   // FIFO
    messages <- "hello"
    messages <- "nothing"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
