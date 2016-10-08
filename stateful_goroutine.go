package main

import (
    "sync"
    "math/rand"
    "sync/atomic"
    "runtime"
    "fmt"
    "time"
)

func main() {
    var state = make(map[int]int)
    var mutex = &sync.Mutex{}
    var ops int64 = 0 //ops will count how many operations we perform against the state


    // 100 reader and 10 writer will manipulate the state map.
    // they use mutex to synchronize the access to the shared state map.

    // simulate 100 reader
    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {
                key := rand.Intn(5)
                mutex.Lock()
                total += state[key]   // use mutex to synchronize access to state
                mutex.Unlock()
                atomic.AddInt64(&ops, 1)
                runtime.Gosched() // explicitly yield the goroutine.
            }
        }()
    }

    // simulate 10 writers
    for w := 0; w < 10; w++ {
        go func() {
            key := rand.Intn(5)
            value := rand.Intn(100)
            mutex.Lock()
            state[key] = value
            mutex.Unlock()
            atomic.AddInt64(&ops, 1)
            runtime.Gosched()
        }()
    }

    time.Sleep(time.Second)

    opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:", opsFinal)

    mutex.Lock()
    fmt.Println("state:", state)
    mutex.Unlock()

}
