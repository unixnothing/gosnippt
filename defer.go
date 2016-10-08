package main

import (
    "fmt"
    "os"
)

//Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup.
// defer is often used where e.g. ensure and finally would be used in other languages
func main() {
    f := createFile("/tmp/defer.txt")
    defer closeFile(f)
    writeFile(f)
}
func writeFile(file *os.File) {
    fmt.Println("writing to file")
    file.Write([]byte("Im nothing,here we go"))
}
func closeFile(file *os.File) {
    fmt.Println("closing file")
    file.Close()
}

func createFile(p string) *os.File {
    fmt.Println("creating file")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}


