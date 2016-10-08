package main

import "fmt"

func main() {
    queue := make(chan string, 5)
    queue <- "a"
    queue <- "b"
    queue <- "c"
    close(queue)

    for ele := range queue{
        fmt.Println(ele)
    }
}
