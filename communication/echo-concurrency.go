package main

import (
    "fmt"
    "net"
    "bufio"
)

func main() {
    listen, err := net.Listen("tcp", "127.0.0.1:8053")
    if err != nil {
        fmt.Printf("Failure to listen %s\n", err.Error())
    }

    for {
        if conn, err := listen.Accept(); err == nil {
            go Echo(conn)
        }
    }
}


func Echo(conn net.Conn) {
    defer conn.Close()
    line, err := bufio.NewReader(conn).ReadString('\n')

    if err != nil {
        fmt.Printf("Failure to read %s\n", err.Error())
        return
    }

    _, err = conn.Write([]byte(line))

    if err != nil {
        fmt.Printf("Failure to write %s\n", err.Error())
        return
    }
}
