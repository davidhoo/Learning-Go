package main

import "fmt"

func fibonacci(i int) (f []int) {
    if i < 2 {
        return
    }
    f = make([]int, i)
    f[0], f[1] = 1, 1
    for x := 2; x < i; x++ {
        f[x] = f[x-1]+f[x-2]
    }
    return f
}

func main() {
    fmt.Println(fibonacci(1))
    fmt.Println(fibonacci(5))
    fmt.Println(fibonacci(10))
}
