// Package zeronl 零基础入门深度学习
package zeronl

// Sensor 是一个神经元结构
type Sensor struct {
	weights   []float32
	bias      float32
	activator func(float32) int8
}

// Init 初始化神经元，设置输入参数的个数，以及激活函数。
// 激活函数的类型为float32 -> float32
func (sensor *Sensor) Init(inputNum int, activator func(float32) int8) {
	sensor.weights = make([]float32, inputNum)
	sensor.bias = 0.0
	sensor.activator = activator
}

// Train 输入训练数据：一组向量、与每个向量对应的label；以及训练轮数、学习率
func (sensor *Sensor) Train(inputVecs [][2]int8, labels []int8, iteration int, rate float32) {
	for i := 0; i < iteration; i++ {
		sensor.oneTrain(inputVecs, labels, rate)
	}
}

// oneTrain 一次迭代，把所有的训练数据过一遍。
func (sensor *Sensor) oneTrain(inputVecs [][2]int8, labels []int8, rate float32) {
	for index, value := range inputVecs {
		// 计算神经元在当前权重下的输出
		output := sensor.Predict(value)
		// 更新权重
		sensor.updateWeights(value, output, labels[index], rate)
	}
}

// updateWeights 按照神经元规则更新权重.
func (sensor *Sensor) updateWeights(inputVec [2]int8, output int8, label int8, rate float32) {
	delta := float32(label - output)
	for index, value := range inputVec {
		sensor.weights[index] = sensor.weights[index] + rate*delta*float32(value)
	}
	sensor.bias += rate * delta
}

// Predict 输入向量，输出神经元的计算结果
func (sensor *Sensor) Predict(inputVec [2]int8) int8 {
	result := sensor.bias
	for index, value := range inputVec {
		result += float32(value) * sensor.weights[index]
	}
	return sensor.activator(result)
}
