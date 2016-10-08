package main

import "fmt"

func main() {
    jobs := make(chan int)
    done := make(chan bool)

    // a go routine for listening the jobs come.
    go func() {
        for {
            j, more := <-jobs  // attention, channel returned 2 values.
            fmt.Println(j, more)
            if !more {
                fmt.Println("no more jobs")
                done <- true
                return // exit the for loop
            }

        }
    }()

    for j := 0; j < 5; j++ {
        fmt.Println("send job:", j)
        jobs <- j
    }

    fmt.Println("all jobs sended, close the jobs channel")
    close(jobs)

    <-done
}
