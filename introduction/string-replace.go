package main

import (
    "fmt"
)

func main() {
    str := "asSASA ddd dsjkdsjs dk"
    b := []byte(str)
    copy(b[4:4+3], []byte("abc"))
    fmt.Println("Before:", str)
    fmt.Println("After :", string(b))
}
