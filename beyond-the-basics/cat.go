package main

import (
    "fmt"
    "os"
    "io"
    "bufio"
    "flag"
)

var nFlag = flag.Bool("n", false, "number each line")

func cat(r *bufio.Reader) {
    i := 1
    for {
        buf, err := r.ReadBytes('\n')

        if err == io.EOF {
            break
        }

        if *nFlag {
            fmt.Fprintf(os.Stdout, "%5d %s", i, buf)
            i++
        } else {
            fmt.Fprintf(os.Stdout, "%s", buf)
        }
    }
    return
}

func main() {
    flag.Parse()

    if flag.NArg() == 0 {
        cat(bufio.NewReader(os.Stdin))
    }

    for i := 0; i < flag.NArg(); i++ {
        f, err := os.Open(flag.Arg(i))
        if err != nil {
            fmt.Fprintf(os.Stdout, "%s: error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
        } else {
            fmt.Fprintf(os.Stdout, "\n%s:\n", flag.Arg(i))
            cat(bufio.NewReader(f))
        }
    }
}
