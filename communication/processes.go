package main

import (
    "fmt"
    "os/exec"
    "sort"
    "strconv"
    "strings"
)

func main() {
    ps := exec.Command("ps", "-e", "-opid,ppid,comm")
    output, _ := ps.Output()
    child := make(map[int][]int)

    for i, s := range strings.Split(string(output), "\n") {
        if i == 0 || len(s) == 0 {
            continue
        }

        fields := strings.Fields(s)
        pid, _ := strconv.Atoi(fields[0])
        ppid, _ := strconv.Atoi(fields[1])
        child[ppid] = append(child[ppid], pid)
    }

    sortChild := make([]int, len(child))
    i := 0
    for k, _ := range child {
        sortChild[i] = k
        i++
    }
    sort.Ints(sortChild)
    for _, ppid := range sortChild {
        fmt.Printf("Pid %d has %d child", ppid, len(child[ppid]))
        if len(child[ppid]) == 1 {
            fmt.Printf(": %v\n", child[ppid])
            continue
        }
        fmt.Printf("ren: %v\n", child[ppid])
    }
}
