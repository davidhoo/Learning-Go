package main

import (
    "fmt"
    "stack"
)

func main() {
    s := new(stack.Stack)
    s.Push(2)
    s.Push(8)
    s.Push(4)
    s.Push(9)
    s.Push(1)
    fmt.Println(s)
    s.Pop()
    s.Pop()
    fmt.Println(s)
}
