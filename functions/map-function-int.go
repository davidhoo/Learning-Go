package main

import "fmt"

func Map(f func(int) int, s []int) []int {
    slice := make([]int, len(s))
    for k, v := range s {
        slice[k] = f(v)
    }
    return slice
}

func main() {
    f := func(i int) int {
        return i + 5
    }

    s := []int{1,2,3,4,5}

    fmt.Println(Map(f, s))
}
