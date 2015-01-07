package main

import "fmt"

func bubbleSort(s []int) {
    for k, v := range s {
        for x := 0; x < len(s); x++ {
            if v < s[x] {
                s[k], s[x] = s[x], s[k]
            }
        }
    }
}

func main() {
    s := []int {1, 5, 2, 6}
    fmt.Println(s)
    bubbleSort(s)
    fmt.Println(s)
}
