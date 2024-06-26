package main

import "fmt"

func main() {
    i := 20
    x := []int{1, 2, 3, 4}
    for i := range(x) {
        fmt.Println("x", i)
    }
    fmt.Printf("i: %v\n", i) // i: 20
}