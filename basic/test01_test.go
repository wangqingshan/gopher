package basic

import (
	"fmt"
	"testing"
)

func Test_go(t *testing.T) {
	go print()
	var b = "this"
	fmt.Println(b)
	c := "this"
	fmt.Println(c)

	var b1 []string
	fmt.Println(b1)
	b2 := make([]string, 0)
	fmt.Println(b2)

	select {}
}

func print() {

	fmt.Println("==")
	x, y := GetUser()
	fmt.Println(x, y)
}

func GetUser() (string, int) {
	return "xiaoming", 19
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestIndex(t *testing.T) {
	var arr1 = [5]int{1, 2, 3, 4, 5}
	var arr2 = [...]int{1, 2}
	fmt.Println(arr2)
	for i := 0; i < len(arr1); i++ {
		fmt.Println("==", arr1[i])
	}

	for _, v := range arr1 {
		fmt.Println("--", v)
	}

	fmt.Println(len(arr1), cap(arr1))
}

func TestSlice(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("打印：", arr)
	// 区间

	fmt.Println("step1:", arr[1:2]) // 前闭后开， 2
	fmt.Println("取所有的元素", arr[:])   // 取所有的元素
	fmt.Println("从某个位置开始", arr[2:])
	fmt.Println("到某个位置结束", arr[:2]) // 下标
}

func TestSliceV2(t *testing.T) {
	var s1 []int
	fmt.Println(s1)

	s2 := []int{2, 3}
	fmt.Println(s2)

	var s3 []int = make([]int, 0)
	fmt.Println(s3)

}

func TestSliceV1(t *testing.T) {
	arr2 := make([]int, 2, 4) // 2 是长度  4是容量  arr2: [0,0]   arr2:=make([]int,2,2)
	arr2[0] = 1               // arr2: [1,0]
	fmt.Println("step1:", arr2)
	arr2 = append(arr2, 3) // append 操作
	arr2 = append(arr2, 3) // append 操作
	fmt.Println("print arr2", arr2)

	//arr3 := []string{"a", "b", "c", "d"}
}

func TestReverseSlice(t *testing.T) {
	arr3 := make([]int, 1, 2)

	arr3[0] = 1

	arr3 = append(arr3, 3)

	// 正序
	for i := 0; i < len(arr3); i++ {
		fmt.Println("正序", arr3[i])
	}

	// 倒序
	for i := 0; i < len(arr3); i++ {
		fmt.Println("倒序：", arr3[len(arr3)-i-1])
	}

	// 倒序2
	for i := len(arr3) - 1; i >= 0; i-- {
		fmt.Println("倒序2：", arr3[i])
	}

}

func TestMapV1(t *testing.T) {
	scoreMap := make(map[string]int, 8) // 初始化，容量8
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)

	if _, ok := scoreMap["张三"]; ok {
		//
		fmt.Println("张三存在")
	}

	val, ok := scoreMap["张三san"]
	fmt.Println(val, "===", ok)

	if _, ok := scoreMap["张三san"]; ok {
		//
		fmt.Println("张三san存在")
	}

	delete(scoreMap, "张三")
	fmt.Println(scoreMap)

	scoreMap["小李"] = 120
	fmt.Println(scoreMap)
}

// func TestXxx(t *testing.T) {

//     var a int

// 	switch  a
// case 0:
// 	// 
// case 0:
// 	// 

// case 0:
// 	// 
// }
