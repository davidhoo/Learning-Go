package main

import "fmt"

func main() {
    p := plustwo()
    fmt.Println(p(3))
}

func plustwo() (func(int) int) {
    return func(x int) int {
        return x + 2
    }
}
