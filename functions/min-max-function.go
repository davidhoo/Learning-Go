package main

import "fmt"

func minmax(s []int) (min int, max int) {
    min,  max = s[0], s[0]
    for _, v := range s {
        if min > v {
            min = v
        }
        if max < v {
            max = v
        }
    }
    return min, max
}

func main() {
    s := []int {1, 3, 6, 2, 9}
    fmt.Println(s)
    fmt.Println(minmax(s))
}
