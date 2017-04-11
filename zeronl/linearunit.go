// Package zeronl 零基础入门深度学习
package zeronl

// LinearUnit 是一个线性单元结构
type LinearUnit struct {
	weights   []float32
	bias      float32
	activator func(float32) float32
}

// Init 初始化线性单元，设置输入参数的个数，以及激活函数。
// 激活函数的类型为float32 -> float32
func (linearUnit *LinearUnit) Init(inputNum int, activator func(float32) float32) {
	linearUnit.weights = make([]float32, inputNum)
	linearUnit.bias = 0.0
	linearUnit.activator = activator
}

// Train 输入训练数据：一组向量、与每个向量对应的label；以及训练轮数、学习率
func (linearUnit *LinearUnit) Train(inputVecs [][1]float32, labels []float32, iteration int, rate float32) {
	for i := 0; i < iteration; i++ {
		linearUnit.oneTrain(inputVecs, labels, rate)
	}
}

// oneTrain 一次迭代，把所有的训练数据过一遍。
func (linearUnit *LinearUnit) oneTrain(inputVecs [][1]float32, labels []float32, rate float32) {
	for index, value := range inputVecs {
		// 计算线性单元在当前权重下的输出
		output := linearUnit.Predict(value)
		// 更新权重
		linearUnit.updateWeights(value, output, labels[index], rate)
	}
}

// updateWeights 按照线性单元规则更新权重.
func (linearUnit *LinearUnit) updateWeights(inputVec [1]float32, output float32, label float32, rate float32) {
	delta := float32(label - output)
	for index, value := range inputVec {
		linearUnit.weights[index] = linearUnit.weights[index] + rate*delta*float32(value)
	}
	linearUnit.bias += rate * delta
}

// Predict 输入向量，输出线性单元的计算结果
func (linearUnit *LinearUnit) Predict(inputVec [1]float32) float32 {
	result := linearUnit.bias
	for index, value := range inputVec {
		result += float32(value) * linearUnit.weights[index]
	}
	return linearUnit.activator(result)
}
