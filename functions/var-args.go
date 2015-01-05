package main

import "fmt"

func printthem(numbers... int) {
    for _, i := range numbers {
        fmt.Println(i)
    }
}

func main() {
    printthem(1,2,3,4,5)
    printthem(22,346,173,3647)
}
