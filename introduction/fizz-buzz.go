package main

import "fmt"

func main() {
    const (
        FIZZ = 3
        BUZZ = 5
    )

    for i := 1; i <= 100; i++ {
        switch {
        case i % (FIZZ * BUZZ) == 0:
            fmt.Println("FizzBuzz")
        case i % FIZZ == 0:
            fmt.Println("Fizz")
        case i % BUZZ == 0:
            fmt.Println("Buzz")
        default:
            fmt.Println(i)
        }
    }
}
