package main

func main() {
    x := "hello"
    y := "nothing"
    x, y = swap(x, y)
    println(x, y)

    m, _ := swap(x, y)          // ignore some return value.
    println(m)
}

func swap(a string, b string) (string, string) {
    return b, a
}
