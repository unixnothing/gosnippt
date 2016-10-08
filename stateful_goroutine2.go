package main

import (
    "math/rand"
    "sync/atomic"
    "time"
    "fmt"
)

type readOp struct {
    key  int
    resp chan int
}

type writeOp struct {
    key  int
    val  int
    resp chan bool
}

// In this example our state will be owned by a single goroutine.
// This will guarantee that the data is never corrupted with concurrent access.
// In order to read or write that state, other goroutines will send messages to the owning goroutine and receive corresponding replies.
func main() {
    var ops int64 = 0

    //The reads and writes channels will be used by other goroutines to issue read and write requests
    reads := make(chan *readOp)
    writes := make(chan *writeOp)

    // Here is the goroutine that owns the state, which is a map as in the previous example but now private to the stateful goroutine.
    go func() {
        var state = make(map[int]int) // is private, just inside this goroutine
        for {
            select {
            case read := <-reads: // every read request for manipulate the state come from the reads channel.
                read.resp <- state[read.key]
            case write := <-writes: // every write request for manipulate the state come from the writes channel.
                state[write.key] = write.val
                write.resp <- true
            }
        }
    }()

    // simulate 100 readers
    for r := 0; r < 100; r++ {
        // This starts 100 goroutines to issue reads to the state-owning goroutine via the reads channel
        go func() {
            for {
                read := &readOp{
                    key:  rand.Intn(5),
                    resp: make(chan int)}
                reads <- read // send read request
                <-read.resp
                atomic.AddInt64(&ops, 1)
            }
        }()
    }

    // simulate 10 writers
    for w := 0; w < 10; w++ {
        go func() {
            write := &writeOp{
                key:rand.Intn(5),
                val:rand.Intn(100),
                resp:make(chan bool)}
            writes <- write
            <-write.resp
            atomic.AddInt64(&ops, 1)
        }()
    }

    time.Sleep(time.Second)
    opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:", opsFinal)
}
