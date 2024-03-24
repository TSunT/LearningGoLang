package main

import "fmt"

func main() {
	myArray := [4]int{1, 2, 3, 4}
	myArray2 := [10]int{1, 2, 3, 4}
	for i := 0; i < len(myArray2); i++ {
		fmt.Println("index = ", i, "value = ", myArray2[i])
	}
	println("------------")
	// _表示匿名变量 只接受不能使用
	pringArray(myArray)
	fmt.Println("------------")
	for _, v := range myArray {
		fmt.Println("value = ", v)
	}
	fmt.Println("-------------")
	// pringArray([4]int(myArray2)) // 编译报错： 数组的长度不匹配
}

func pringArray(myArray [4]int) {
	// _表示匿名变量 只接受不能使用
	for index, v := range myArray {
		fmt.Println("index = ", index, "value = ", v)
	}
	myArray[0] = 100
}
