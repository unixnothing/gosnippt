package main

import (
    "fmt"
    "time"
)

func worker(finished chan bool) {
    fmt.Println("Worker: Started")
    time.Sleep(time.Second)
    fmt.Println("Worker: Finished")
    finished <- true  // worker tell others: i'm finished.  Use the chan to synchronize the gorutines
}

func main() {
    finished := make(chan bool)

    fmt.Println("Main: Starting worker")

    go worker(finished)

    go func(message string) {
        fmt.Println("Im in gorutine", message)
    }("Tom")

    fmt.Println("Main: Waiting for worker to finish")
    <-finished
    fmt.Println("Main: Completed")
}
