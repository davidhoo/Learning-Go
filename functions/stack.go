package main

import "fmt"

type stack struct {
    i    int
    data [10]int
}

func (s *stack) push(v int) {
    s.data[s.i] = v
    s.i++
}

func (s *stack) pop() int {
    if s.i == 0 {
        return 0
    }
    s.i--
    return s.data[s.i]
}

func (s *stack) String() string {
    var str string
    for i := 0; i < s.i; i++ {
        str += fmt.Sprintf("[%v:%v] ", i, s.data[i])
    }
    return str
}

func main() {
    s := new(stack)
    s.push(1)
    s.push(2)
    s.push(3)
    s.push(4)
    s.push(7)
    s.push(1)
    s.push(0)
    fmt.Printf("My stact %v\n", s)
    fmt.Println(s.pop())
    fmt.Printf("My stact %v\n", s)
    fmt.Println(s.pop())
    fmt.Println(s.pop())
    fmt.Println(s.pop())
    fmt.Println(s.pop())
    fmt.Printf("My stact %v\n", s)
}
