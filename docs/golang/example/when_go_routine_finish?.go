package main

import "fmt"

func main() {
	fmt.Println("start")
	for i:=range 10{
		go fmt.Println(i) 
	}
	fmt.Println("end")
}