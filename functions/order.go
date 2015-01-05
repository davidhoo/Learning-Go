package main

import "fmt"

func order(x, y int) (int, int) {
    if x > y {
        return y, x
    }

    return x, y
}

func main() {
    fmt.Println(order(2, 7))
    fmt.Println(order(7, 2))
}
