package main

import "fmt"

func main() {
    slice := []float64{2.0, 3.0, 4.0, 5.0,6.0}
    var sum float64

    for _, f := range slice {
        sum += f
    }

    average := sum / float64(len(slice))

    fmt.Printf("slice %v average is %f\n", slice, average)
}
