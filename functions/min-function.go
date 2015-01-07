package main

import "fmt"

func min(s []int) int {
    min := s[0]
    for _, v := range s {
        if min > v {
            min = v
        }
    }
    return min
}

func main() {
    s := []int {1, 5, 2, 8}
    fmt.Println(s)
    fmt.Println(min(s))
}
