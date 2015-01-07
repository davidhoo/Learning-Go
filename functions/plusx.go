package main

import "fmt"

func plusx(x int) (func(int) int) {
    return func(i int) int {
        return i + x
    }
}

func main() {
    p := plusx(3)
    fmt.Println(p(2))
}
