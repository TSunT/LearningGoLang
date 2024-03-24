package main

import "fmt"

func main() {
	var mySlice1 []int // []内不指定长度则是切片, 注意，此时没有分配内存空间！
	fmt.Println("slice2 length:", len(mySlice1), "cap =", cap(mySlice1), " value: ", mySlice1)
	mySlice1 = []int{1, 2, 3, 4} // 开辟空间
	printSilce(mySlice1)
	for i, v := range mySlice1 {
		fmt.Println("index = ", i, "value = ", v)
	}
	mySlice2 := make([]int, 10, 11) // 声明一个切片 ，并分配（len） 10个int元素空间 初始化都是默认值,
	// 第3个参数是cap 表示当前分配的容量，如果后续append没有超过其容量则不用去分配 如果超过cap 则再分配当前cap个空间
	// fmt.Println("slice2 length:", len(mySlice2), "cap =", cap(mySlice2), " value: ", mySlice2)
	printSilce(mySlice2)
	fmt.Println("append to slice2...")
	mySlice2 = append(mySlice2, 1)
	pringSliceInfo(mySlice2)
	fmt.Println("append to slice2...")
	mySlice2 = append(mySlice2, 2)
	pringSliceInfo(mySlice2)
	for i, value := range mySlice2 {
		fmt.Println("index = ", i, "value = ", value)
	}
}

func printSilce(mySlice []int) {
	// 切片是引用传递 改变形参的值 实参也会改变
	println("----- intro func -----")
	for i := 0; i < len(mySlice); i++ {
		fmt.Println("index = ", i, "value = ", mySlice[i])
	}
	mySlice[0] = 1000
	println("----- outro func -----")
}

func pringSliceInfo(mySlice []int) {
	fmt.Println("slice2 length:", len(mySlice), "cap =", cap(mySlice), " value: ", mySlice)
}
