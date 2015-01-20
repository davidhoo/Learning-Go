package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "io"
)

func main() {
    var chars, words, lines int
    reader := bufio.NewReader(os.Stdin)
    for {
        str, ok := reader.ReadString('\n');
        if ok == io.EOF {
            return
        }
        chars += len(str)
        words += len(strings.Fields(str))
        lines++
        fmt.Printf("str:%schars:%d words:%d lines:%d\n", str, chars, words, lines)
    }
}
