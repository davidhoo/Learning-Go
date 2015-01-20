package main

import (
    "fmt"
    "runtime"
    "os"
    "bufio"
)

func main() {
    _, fullFileName, _, _ := runtime.Caller(0)
    f, _ := os.Open(fullFileName)
    reader := bufio.NewReader(f)
    for {
        str, e := reader.ReadString('\n')
        if e != nil {
            return
        }
        fmt.Printf("%s", str)
    }
}
