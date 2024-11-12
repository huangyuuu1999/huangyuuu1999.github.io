package main
// import "fmt"

func main() {
  // 两个map的对比？
  a := map[int]int{1:2, 4:9}
  b := map[int]int{4:9, 1:2}
  // fmt.Println(a==b) // invalid operation: a == b (map can only be compared to nil)
  _, _ = a, b
}
