package main

import "fmt"

var (
	g_num int     = 112
	g_fl  float64 = 90.782
	g_str string  = "2xx"
)

// iota 只能在cost()里面使用
const (
	a, b = iota + 1, iota + 2 // iota = 0
	c, d                      // iota = 1 公式自动为上一行 其中iota从0开始每行累加 1
	e, f = iota * 3, iota * 4 // iota = 2
	g, h                      // iota = 3
)

func main() {
	var num int = 18
	str := "dds"
	fmt.Println("num = ", num)
	fmt.Println("str value := ", str)
	const NEE int = 1
	num2 := num * NEE
	fmt.Println("num2 =", num2)

	fmt.Printf("a= %d, b= %d, c= %d, d= %d, e= %d, f= %d, g= %d, h= %d", a, b, c, d, e, f, g, h)

	const (
		a1 = 1
		b1 = 2
	)

}
