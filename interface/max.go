package main

import "fmt"


func max(s interface{}) interface{} {
    switch s.(type) {
    case []int:
        max := s.([]int)[0]
        for _, v := range s.([]int) {
            if v > max {
                max = v
            }
        }
        return max
    case []float32:
        max := s.([]float32)[0]
        for _, v := range s.([]float32) {
            if v > max {
                max = v
            }
        }
        return max
    }
    return 0
}

func less(x, y interface{}) bool {
    switch x.(type) {
    case int:
        if _, ok := x.(int); ok {
            return x.(int) < y.(int)
        }
    case float32:
        if _, ok := x.(float32); ok {
            return x.(float32) < y.(float32)
        }
    }
    return false
}

func main() {
    s1 := []float32{1.0, 2.2, 3.1, 5.1, 2.4}
    s2 := []int{1, 4, 6, 2, 9}
    fmt.Printf("s1:%v\nmax:%v\n", s1, max(s1))
    fmt.Printf("s2:%v\nmax:%v\n", s2, max(s2))
    x, y, z := 3, 4, 0
    if z = x; less(x, y) {
        z = y
    }
    fmt.Printf("x:%v\ny:%v\nz:%v\n", x, y, z)
}
