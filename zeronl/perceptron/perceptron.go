// Package perceptron 是感知器相关内容
package perceptron

// Perceptron 是一个感知器结构
type Perceptron struct {
	weights   []float32
	bias      float32
	activator func(float32) int8
}

// Init 初始化感知器，设置输入参数的个数，以及激活函数。
// 激活函数的类型为float32 -> float32
func (perceptron *Perceptron) Init(inputNum int, activator func(float32) int8) {
	perceptron.weights = make([]float32, inputNum)
	perceptron.bias = 0.0
	perceptron.activator = activator
}

// Train 输入训练数据：一组向量、与每个向量对应的label；以及训练轮数、学习率
func (perceptron *Perceptron) Train(inputVecs [][2]int8, labels []int8, iteration int, rate float32) {
	for i := 0; i < iteration; i++ {
		perceptron.oneTrain(inputVecs, labels, rate)
	}
}

// oneTrain 一次迭代，把所有的训练数据过一遍。
func (perceptron *Perceptron) oneTrain(inputVecs [][2]int8, labels []int8, rate float32) {
	for index, value := range inputVecs {
		// 计算感知器在当前权重下的输出
		output := perceptron.Predict(value)
		// 更新权重
		perceptron.updateWeights(value, output, labels[index], rate)
	}
}

// updateWeights 按照感知器规则更新权重.
func (perceptron *Perceptron) updateWeights(inputVec [2]int8, output int8, label int8, rate float32) {
	delta := float32(label - output)
	for index, value := range inputVec {
		perceptron.weights[index] = perceptron.weights[index] + rate*delta*float32(value)
	}
	perceptron.bias += rate * delta
}

// Predict 输入向量，输出感知器的计算结果
func (perceptron *Perceptron) Predict(inputVec [2]int8) int8 {
	result := perceptron.bias
	for index, value := range inputVec {
		result += float32(value) * perceptron.weights[index]
	}
	return perceptron.activator(result)
}
