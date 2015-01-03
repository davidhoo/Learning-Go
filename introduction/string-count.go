package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    str := "asSASA ddd dsjkdsjs dk"
    b := []byte(str)
    fmt.Printf("String %s Length: %d Runes: %d\n", str, len(b), utf8.RuneCountInString(str))
}
