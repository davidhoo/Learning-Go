package main

import "fmt"

type T interface{}

func Map(f func(T) T, s []T) []T {
    slice := make([]T, len(s))
    for k, v := range s {
        slice[k] = f(v)
    }
    return slice
}

func main() {
    f := func(t T) T {
        switch t.(type) {
        case int:
            return t.(int) * 2
        case string:
            return t.(string) + t.(string)
        }
        return t
    }

    s1 := []T{1,2,3,4,5}
    fmt.Println(Map(f, s1))
    s2 := []T{"a", "b", "c", "d"}
    fmt.Println(Map(f, s2))
}
