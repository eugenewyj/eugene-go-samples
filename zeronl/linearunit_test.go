package zeronl

import (
	"fmt"
	"testing"
)

// 定义线性单元激活函数
func linearUnitActivator(x float32) float32 {
	return x
}

// 虚拟5个人的工作年限及收入
func getLinearUnitTrainingDataSet() ([][1]float32, []float32) {
	// 构建训练数据
	// 输入向量列表， 每一项是工作年限
	inputVecs := [][1]float32{{5.0}, {3.0}, {8.0}, {1.4}, {10.1}}
	// 期望的输出列表，月薪，注意要与输入一一对应
	labels := []float32{5500.0, 2300.0, 7600.0, 1800.0, 11400.0}
	return inputVecs, labels
}

// 测试感知器
func TestLinearUnit(t *testing.T) {
	// 创建感知器，输入参数个数为2（因为and是二元函数），激活函数为f
	linearUnit := &LinearUnit{}
	linearUnit.Init(1, linearUnitActivator)

	// 训练，迭代10轮, 学习速率为0.1
	inputVecs, labels := getLinearUnitTrainingDataSet()
	linearUnit.Train(inputVecs, labels, 10, 0.01)

	// 打印训练获得的权重
	fmt.Printf("%+v\n", linearUnit)
	// 测试
	fmt.Printf("Work 3.4 year, monthly salary = %.2f\n", linearUnit.Predict([1]float32{3.4}))
	fmt.Printf("Work 15 year, monthly salary = %.2f\n", linearUnit.Predict([1]float32{15}))
	fmt.Printf("Work 1.5 year, monthly salary = %.2f\n", linearUnit.Predict([1]float32{1.5}))
	fmt.Printf("Work 6.3 year, monthly salary = %.2f\n", linearUnit.Predict([1]float32{6.3}))
}
