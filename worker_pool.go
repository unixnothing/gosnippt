package main

import (
    "fmt"
    "time"
    "strconv"
)

// worker take job off FROM jobs chan, and set result INTO rest chan
func myworker(workId int, jobs <-chan int, rest chan <- string) {
    for j := range jobs {
        fmt.Printf("worker[%d] take the job:%d \n", workId, j)
        time.Sleep(time.Second)
        rest <- "job: " + strconv.Itoa(j) + " done"
    }
}

func main() {
    jobs := make(chan int, 10)
    result := make(chan string, 10)

    // init all jobs
    for j := 0; j < 10; j++ {
        jobs <- j
    }
    // need to close jobs chann to prevent the myworker 'range jobs' dead lock
    close(jobs)

    go myworker(1, jobs, result)
    go myworker(2, jobs, result)

    for i := 0; i < 10; i++ {
        r := <-result
        fmt.Println(r)
    }

    //for rest := range result {
    //    fmt.Println(rest)
    //}
}
