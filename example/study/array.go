package main

import "fmt"

// 数组使用
func main() {
	// 数组声明
	var arr1 [10]string
	var arr2 = [10]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	var arr3 = []string{}
	var arr4 = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	// 5行2列
	var a1 = [5][4]int{}

	// n行4列
	var a2 = [][4]int{{0, 1, 2, 3}, {4, 5, 6, 7}}

	// n行n列
	var a3 [][]string
	//
	arr1[1] = "hi"
	arr1[2] = "hello"
	fmt.Println(arr1, arr1[1], arr2, arr3, arr4)

	//
	a1[2][1] = 66
	a1[1][0] = 66
	a1[1][1] = 66
	fmt.Println(a1[2])
	fmt.Println(a1)
	//
	a2[0][1] = 6666
	fmt.Println(a2[1])
	fmt.Println(a2, a3)
}
