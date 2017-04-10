// Package perceptron 是感知器相关内容
package perceptron

// Perceptron 是一个感知器结构
type Perceptron struct {
	weights []float32
	bias    float32
	fn      func(float32) float32
}

// Init 初始化感知器，设置输入参数的个数，以及激活函数。
// 激活函数的类型为float32 -> float32
func (perceptron *Perceptron) Init(inputNum int, f func(float32) float32) {
	perceptron.weights = make([]float32, inputNum)
	perceptron.fn = f
}

// Train 输入训练数据：一组向量、与每个向量对应的label；以及训练轮数、学习率
func (perceptron *Perceptron) Train(inputVecs [][2]int8, labels []int8, iteration int, rate float64) {

}

// Predict 输入向量，输出感知器的计算结果
func (perceptron *Perceptron) Predict(inputVec [2]int8) int8 {
	return 1
}
