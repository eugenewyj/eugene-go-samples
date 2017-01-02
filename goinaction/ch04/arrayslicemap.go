package main

import (
	"fmt"
)

func main() {
	arrayTest()
	sliceTest()
	mapTest()
}
func arrayTest()  {
	// 数组相关代码演示
	var array1 [5]int
	for i := 0; i < len(array1); i++ {
		fmt.Printf("array1[%d] is %d \n", i, array1[i])
	}

	array2 := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < len(array2); i++ {
		fmt.Printf("array2[%d] is %d \n", i, array2[i])
	}

	array3 := [5]int{1:10, 2: 20}
	for i := 0; i < len(array3); i++ {
		fmt.Printf("array3[%d] is %d \n", i, array3[i])
	}

	array2[2] = 35
	for i := 0; i < len(array2); i++ {
		fmt.Printf("array2[%d] is %d \n", i, array2[i])
	}

	array4 := [5]*int{0: new(int), 1: new(int)}
	temp1 := 10
	array4[2] = &temp1
	array4[3] = new(int)
	array4[4] = new(int)
	*array4[3] = 30
	for i := 0; i < len(array4); i++ {
		fmt.Printf("array4[%d] is %d \n", i, *array4[i])
	}

	var array21 [5]string
	array22 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	array21 = array22
	for i := 0; i < len(array21); i++ {
		fmt.Printf("array4[%d] is %s \n", i, array21[i])
	}

	//不同长度的同类型数据不能互相赋值，编译错误
	//var array23 [4]string
	//array23 = array22

	var array31 [3]*string
	array32 := [3]*string{new(string), new(string), new(string)}
	*array32[0] = "Red"
	*array32[1] = "Blue"
	*array32[2] = "Green"
	array31 = array32
	for i := 0; i < len(array31); i++ {
		fmt.Printf("array31[%d] is %d value is %s \n", i, array31[i], *array31[i])
	}
	for i := 0; i < len(array32); i++ {
		fmt.Printf("array32[%d] is %d value is %s \n", i, array32[i], *array32[i])
	}

	//多维数组
	var arrays1 [4][2]int
	fmt.Println(arrays1)
	arrays2 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println(arrays2)
	arrays3 := [4][2]int{1: {20, 21}, 3: {40, 41}}
	fmt.Println(arrays3)
	arrays4 := [4][2]int{1: {0: 20}, 3: {1: 40}}
	fmt.Println(arrays4)
	var arrays5 [2][2]int
	arrays5[0][0] = 10
	arrays5[0][1] = 20
	arrays5[1][0] = 30
	arrays5[1][1] = 40
	var arrays6 [2][2]int
	arrays6 = arrays5
	fmt.Println(arrays6)

	var arrays7 [2]int = arrays5[1]
	fmt.Println(arrays7)
	var val int = arrays5[1][0]
	fmt.Println(val)

	// 参数传递时数组是按值传递，并且是深拷贝。
}
// 数组相关操作演示
func sliceTest() {
	//长度和容量都是5
 	slice1 := make([]int, 5)
	fmt.Println("slice1:", slice1)

	//长度为3，容量为5
	slice2 := make([]int, 3, 5)
	fmt.Println("slice2:", slice2)

	// 长度大于容量编译错误
	//slice3 := make([]int, 5, 3)

	// 声明slice并初始化
	slice4 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Println("slice4", slice4)
	slice5 := []int{10, 20, 30}
	fmt.Println("slice5", slice5)

	// 采用索引位置声明slice
	slice6 := []int{10: 10}
	fmt.Println("slice6", slice6)

	// nil silce
	var slice7 []int
	fmt.Println("slice7", slice7)

	// empty sliceTest
	slice8 := make([]int, 0)
	fmt.Println("slice8", slice8)
	slice9 := []int{}
	fmt.Println("slice9", slice9)

	// change value
	slice10 := []int{10, 20, 30, 40, 50}
	slice10[1] = 5
	fmt.Println("slice10", slice10)

	// taking the sliceTest of a sliceTest
	// sliceTest[i:j]
	// length: j - i
	// capacity: k - i  k为原有slice的capacity
	slice11 := []int{10, 20, 30, 40, 50}
	newSlice1 := slice11[1:3]
	newSlice1[1] = 35
	newSlice1 = append(newSlice1, 60)
	fmt.Println("slice11:", slice11)
	fmt.Println("newSlice1:", newSlice1)
	newSlice1 = append(newSlice1, 70)
	fmt.Println("slice11:", slice11)
	fmt.Println("newSlice1:", newSlice1)
	newSlice1 = append(newSlice1, 80)
	fmt.Println("slice11:", slice11)
	fmt.Println("newSlice1:", newSlice1)

	slice12 := []int{10, 20, 30, 40}
	newSlice2 := append(slice12, 50)
	fmt.Println("slice12:", slice12)
	fmt.Println("newSlice2:", newSlice2)

	// three index slices
	// For sliceTest[i:j:k]
	// length: j - i
	// Capacity: k - i
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	target := source[2:3:4]
	fmt.Println("source:", source)
	fmt.Println("target:", target)

	target2 := source[2:3:3]
	fmt.Println("source:", source)
	fmt.Println("target2:", target2)
	target2 = append(target2, "Kiwi")
	fmt.Println("source:", source)
	fmt.Println("target2:", target2)

	s1 := []int{1, 2}
	s2 := []int{3, 4}
	fmt.Printf("%v\n", append(s1, s2...))
	for index, value := range s1 {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}
}

func	 mapTest() {
	dict1 := make(map[string]int)
	fmt.Println("dict1:", dict1)
	dict2 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	fmt.Println("dict2:", dict2)

	// slice不能作为map的key,但是可以为value
	//编译错误
	//dict3 := make(map[[]string]int)
	dict4 := make(map[int][]string)
	fmt.Println("dict4:", dict4)

	dict5 := map[string]string{}
	dict5["Red"] = "#da1337"
	fmt.Println("dict5:", dict5)

	//赋值给nil map，运行时异常
	//var dict6 map[string]string
	//dict6["Red"] = "#da1337"
	//fmt.Println("dict6:", dict6)

	//获取值
	value, exists := dict2["Red"]
	if exists {
		fmt.Println("dict2[Red]= ", value)
	}

	value2 := dict2["Orange"]
	if value2 != "" {
		fmt.Println("dict2[Orange]=", value2)
	}

	for key, value := range dict2 {
		fmt.Printf("Key: %s Value: %s \n", key, value)
	}

	delete(dict2, "Red")
	for key, value := range dict2 {
		fmt.Printf("Key: %s Value: %s \n", key, value)
	}

	// map 作为参数传递时传递的是地址
	dict6 := map[string]string{ "AliceBlue": "#f0f8ff", "Coral": "#ff7F50", "DarkGray": "#a9a9a9", "ForestGreen": "#228b22", }
	for key, value := range dict6 {
		fmt.Printf("remove before dict6 Key: %s Value: %s \n", key, value)
	}
	removeMap(dict6, "Coral")
	for key, value := range dict6 {
		fmt.Printf("remove after dict6 Key: %s Value: %s \n", key, value)
	}
}

func removeMap(dict map[string]string, value string) {
	delete(dict, value)
}
