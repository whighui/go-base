// main.go
package main

var nodeLinkInfo map[int32][]int32  //key代表点 value代表边数组
var linkNodeInfo map[int32][2]int32 //key代表边，value[0]是边的起点ID value[1]是边的终点ID

type Node struct {
	nodeId int32   //nodeID代表地图中的一个节点
	linkId int32   //linkID代表某个节点上的一条边
	next   []*Node //上述点和边组成的路段的临近路况信息
}

// 这个是代表司机就是正在这个  A--司机---B 司机在当前边上和起点A 想要展示地图1000范围内的全貌
// 假如传进来的root节点是已知rootID和linkID
func generate(root *Node) *Node {
	if root == nil || root.nodeId == 0 {
		return nil
	}
	root.next = make([]*Node, 0)
	linkList := nodeLinkInfo[root.nodeId]
	if len(linkList) == 0 {
		return nil
	}
	for _, linkID := range linkList {
		nodeArray := linkNodeInfo[linkID]
		if nodeArray[1] == 0 { //也即就是这条边就是没有终点的呗
			return nil
		}
		for _, nodeID := range nodeArray {
			node := &Node{
				nodeId: nodeID,
				linkId: linkID,
			}
			root.next = append(root.next, node)

			generate(node) //这里边面试官告诉我先
		}
	}
	return root
}
