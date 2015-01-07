package main

import "fmt"

func max(s []int) int {
    max := s[0]
    for _, v := range s {
        if v < max {
            max = v
        }
    }
    return max
}

func main() {
    s := []int{1, 4, 5, 6, 3, 58, 32}
    fmt.Println(s)
    fmt.Println(max(s))
}
