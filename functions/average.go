package main

import "fmt"

func average(slice []float64) (avg float64) {
    if len(slice) == 0 {
        return
    }

    sum := 0.0
    for _, f := range slice {
        sum += f
    }

    return sum / float64(len(slice))
}

func main() {
    slice := []float64{2.0, 3.0, 4.0, 5.0, 6.0}
    fmt.Println(average(slice))
}
