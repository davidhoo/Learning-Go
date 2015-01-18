package main

import "fmt"

func main() {
    ch := make(chan int)
    q := make(chan bool)
    go shower(ch, q)

    for i := 0; i < 10; i++ {
        ch <- i
    }
    q<-true
}

func shower(ch chan int, q chan bool) {
    for {
        select {
        case i := <- ch:
            fmt.Println(i)
        case <-q:
            break
        }
    }
}
