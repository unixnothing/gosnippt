package main

import (
    "errors"
    "fmt"
)

func main() {
    fmt.Println(addme(1, 3))
    fmt.Println(addme(1, -3))

    fmt.Println(multiple(1, 3))
    fmt.Println(multiple(1, -3))

    ret, e := multiple(1, -3)
    if e != nil {
        fmt.Println("multiple not work, message:", e)
    } else {
        fmt.Println(ret)
    }
}

func addme(x int, y int) (int, error) {
    if x < 0 || y < 0 {
        return 0, errors.New("Cann't work with negative!")
    }
    return x + y, nil
}

type argError struct {
    arg     []int
    message string
}

// argError impl the Error method, so it can be used as error
func (e *argError) Error() string {
    return fmt.Sprintf("%d ---- %s", e.arg, e.message)
}

func multiple(x int, y int) (int, error) {
    if x < 0 || y < 0 {
        return 0, &argError{[]int{x, y}, "not postive"}
    }
    return x * y, nil
}
