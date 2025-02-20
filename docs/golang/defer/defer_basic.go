// filename defer_basic.go
// show basic usage of defer

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// deferCloseFile()
	// deferRecover()
	deferLog()
}

func deferCloseFile() {
	file, _ := os.Open("example.txt")
	defer file.Close()
	data := make([]byte, 20)
	n, _ := file.Read(data)
	fmt.Println("Data read from file:\n", string(data[:n]))
	fmt.Printf("\n%v\n", "other code...")
}

func mightPanic() {
	panic("a problem")
}

func deferRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("revovered a panic: %v\n", r)
		}
	}()
	mightPanic()
	fmt.Println("after panic.")
}

func deferLog() {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		fmt.Printf("Function took: %v\n", duration)
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("Function completed")
}