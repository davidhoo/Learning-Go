package main

import "fmt"

func Map(f func(s string) string, s []string) []string {
    slice := make([]string, len(s))
    for k, v := range s {
        slice[k] = f(v)
    }
    return slice
}

func main() {
    s := []string{"a", "b", "c", "d"}
    f := func(s string) string {
        return s + s
    }

    fmt.Println(Map(f, s))
}
