package main

import (
	"fmt"

	"github.com/eugenewyj/go-sample/zeronl/perceptron"
)

// 定义激活函数
func f(x float32) float32 {
	var result float32
	if x > 0 {
		result = 1.0
	} else {
		result = 0.0
	}
	return result
}

// 使用and真值表训练感知器
func trainAndPerceptron() *perceptron.Perceptron {
	// 创建感知器，输入参数个数为2（因为and是二元函数），激活函数为f
	p := &perceptron.Perceptron{}
	p.Init(2, f)

	// 训练，迭代10轮, 学习速率为0.1
	inputVecs, labels := getTrainingDataSet()
	p.Train(inputVecs, labels, 10, 0.1)

	// 返回训练好的感知器
	return p

}

// 基于and真值表构建训练数据
func getTrainingDataSet() ([][2]int8, []int8) {
	inputVecs := [][2]int8{{1, 1}, {1, 0}, {0, 1}, {0, 0}}
	labels := []int8{1, 0, 0, 0}
	return inputVecs, labels
}

func main() {
	// 训练and感知器
	andPerceptron := trainAndPerceptron()
	// 打印训练获得的权重
	fmt.Printf("%+v\n", andPerceptron)
	// 测试
	fmt.Printf("1 and 1 = %d\n", andPerceptron.Predict([2]int8{1, 1}))
	fmt.Printf("1 and 0 = %d\n", andPerceptron.Predict([2]int8{1, 0}))
	fmt.Printf("0 and 1 = %d\n", andPerceptron.Predict([2]int8{0, 1}))
	fmt.Printf("0 and 0 = %d\n", andPerceptron.Predict([2]int8{0, 0}))
}
