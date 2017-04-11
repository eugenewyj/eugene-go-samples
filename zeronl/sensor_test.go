package zeronl

import (
	"fmt"
	"testing"
)

// 定义感知器激活函数
func sensorActivator(x float32) int8 {
	var result int8
	if x > 0 {
		result = 1
	} else {
		result = 0
	}
	return result
}

// 基于and真值表构建训练数据
func getAndSensorTrainingDataSet() ([][2]int8, []int8) {
	inputVecs := [][2]int8{{1, 1}, {1, 0}, {0, 1}, {0, 0}}
	labels := []int8{1, 0, 0, 0}
	return inputVecs, labels
}

// 测试感知器
func TestSensor(t *testing.T) {
	// 创建感知器，输入参数个数为2（因为and是二元函数），激活函数为f
	sensor := &Sensor{}
	sensor.Init(2, sensorActivator)

	// 训练，迭代10轮, 学习速率为0.1
	inputVecs, labels := getAndSensorTrainingDataSet()
	sensor.Train(inputVecs, labels, 10, 0.1)

	// 打印训练获得的权重
	fmt.Printf("%+v\n", sensor)
	// 测试
	fmt.Printf("1 and 1 = %d\n", sensor.Predict([2]int8{1, 1}))
	fmt.Printf("1 and 0 = %d\n", sensor.Predict([2]int8{1, 0}))
	fmt.Printf("0 and 1 = %d\n", sensor.Predict([2]int8{0, 1}))
	fmt.Printf("0 and 0 = %d\n", sensor.Predict([2]int8{0, 0}))
}
