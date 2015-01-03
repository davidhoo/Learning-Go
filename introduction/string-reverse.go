package main

import (
    "fmt"
)

func main() {
    str := "asSASA ddd dsjkdsjs dk"
    fmt.Printf("Before:%s\n", str)

    b := []byte(str)
    fmt.Print("After :")
    for i := len(b) - 1; i >= 0; i-- {
        fmt.Printf("%s", string(b[i]))
    }
    fmt.Println()
}
