package main

import "fmt"

func main() {
    var array [10]int
    for k, _ := range array {
        array[k] = k
    }
    fmt.Println(array)
}
