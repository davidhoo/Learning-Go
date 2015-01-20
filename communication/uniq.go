package main

import (
    "fmt"
)

func main() {
    arr := []string{"a", "b", "c", "c", "d", "e", "e", "f"}
    fmt.Println(arr)
    first := arr[0]
    fmt.Printf("%s ", first)
    for _, s := range arr {
        if s != first {
            first = s
            fmt.Printf("%s ", s)
        }
    }
    fmt.Println()
}
