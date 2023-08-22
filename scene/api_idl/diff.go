package idl

// DiffIdlAnalyse 当判断有新增节点 那么就是判断存在差异 输出为新增节点  检测json必须先要转化为IDL 在进行做判断 如果边转边判断 那么就是代码耦合太深 这种形式就是不好
// 1.检测 key 出现新的节点
// 2.检测 key 属性变化
func DiffIdlAnalyse(baseIdl, detectedIdl *TreeNode) bool { //如果存在diff 返回true
	//存储
	curBaseList, curDetectedList, nextBaseList, nextDetectedList := make([]*TreeNode, 0), make([]*TreeNode, 0), make([]*TreeNode, 0), make([]*TreeNode, 0)

	curBaseList, curDetectedList = append(curBaseList, baseIdl), append(curDetectedList, detectedIdl)

	for len(curBaseList) != 0 && len(curDetectedList) != 0 {

		if judgeLayerNodeIsDiff(curBaseList, curDetectedList) { //如果检测比原始idl 出现新的节点返回true 证明存在新的diff
			return true
		}
		for _, node := range curBaseList {
			if node.Child != nil {
				nextBaseList = append(nextBaseList, node.Child...)
			}
		}
		for _, node := range curDetectedList {
			if node.Child != nil {
				nextDetectedList = append(nextDetectedList, node.Child...)
			}
		}

		curBaseList, curDetectedList = nextBaseList, nextDetectedList
		nextBaseList, nextDetectedList = make([]*TreeNode, 0), make([]*TreeNode, 0)
	}

	if len(curBaseList) > 0 && len(curDetectedList) == 0 {
		return false
	} else if len(curBaseList) == 0 && len(curDetectedList) == 0 { //如果张的一样不存在diff
		return false
	}
	return true //检测list出现新的节点。那么就是直接返回true 判断存在diff
}

// judgeLayerNodeIsDiff
// 1.判断层 检测Node相比于baseNode是否出现新的节点  出现新的节点就是返回true
func judgeLayerNodeIsDiff(baseList, detectedList []*TreeNode) bool {
	for _, detectedNode := range detectedList {
		flag := false
		for _, baseNode := range baseList {
			if *(detectedNode.Info) == *(baseNode.Info) {
				flag = true
				break
			}
		}
		if !flag { //如果原始里边没有 那么就是返回ture 存在diff
			return true
		}
	}
	return false
}
