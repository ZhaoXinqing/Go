package changjing_ai

import "fmt"

// main ... 合并切片
func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4, 5, 6}
	a = append(a, b...)
	fmt.Println(a)
}
