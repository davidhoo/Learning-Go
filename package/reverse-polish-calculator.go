package main

import (
    "fmt"
    "stack"
    "bufio"
    "os"
    "strconv"
)

func main() {
    st := new(stack.Stack)
    reader := bufio.NewReader(os.Stdin)
    var token string
    for {
        s, err := reader.ReadString('\n')
        if err != nil {
            return
        }

        for _, c := range s {
            switch {
            case c >= '0' && c <= '9':
                token += string(c)
            case c == ' ':
                r, _ := strconv.Atoi(token)
                st.Push(r)
                token = ""
            case c == '+':
                y, x := st.Pop(), st.Pop()
                fmt.Printf("%v+%v=%v\n", x, y, x+y)
            case c == '*':
                x, y := st.Pop(), st.Pop()
                fmt.Printf("%v*%v=%v\n", x, y, x*y)
            case c == '-':
                x, y := st.Pop(), st.Pop()
                fmt.Printf("%v-%v=%v\n", x, y, x-y)
            case c == '/':
                x, y := st.Pop(), st.Pop()
                fmt.Printf("%v/%v=%v\n", x, y, x/y)
            case c == 'q':
                return
            }
        }
    }
}
