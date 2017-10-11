package zeronl

// Connection 连接对象，主要负责记录连接的权重，以及这个连接所关联的上下游节点。
type Connection struct {
	upstreamNode   *Node
	downstreamNode *Node
	weight         float32
	gradient       float32
}

// Node 神经网络节点， 负责记录和维护节点自身信息及与这个节点关联的上下游连接，实际输出值和误差项计算。
type Node struct {
	layerIndex int8
	nodeIndex  int8
	output     float32
	delta      float32
	downstream []*Connection
	upstream   []*Connection
}

// SetOutput 设置节点输出值。如果节点属于输入层会用的这个函数。
func (node *Node) SetOutput(output float32)  {
	node.output = output
}

// AppendDownstreamConnection 添加一个到下游节点的连接。
func (node *Node) AppendDownstreamConnection(connection *Connection) {
	node.downstream = append(node.downstream, connection)
}

// AppendUpstreamConnection 添加一个到上游节点的连接。
func (node *Node) AppendUpstreamConnection(connection *Connection)  {
	node.upstream = append(node.upstream, connection)
}

