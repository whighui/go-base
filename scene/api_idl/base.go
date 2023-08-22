package idl

type TreeNode struct {
	Info  *NodeInfo
	Child []*TreeNode
}

type NodeInfo struct {
	Tag       string
	ParamKey  string
	ParamType string
	Layer     int
}
