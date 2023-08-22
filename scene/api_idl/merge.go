package idl

import (
	"context"
	"sort"
)

// MergeHistoryIdl
// Merge标准    出现新节点在对应层级下进行补充
//
//	节点在样本中出现过，判断孩子是否存在差异，如果存在差异进行补充
//
// sampleIdl  第一份样本IDL  historyIdl 其他历史IDL
// 如果出现变更 进行update
func MergeHistoryIdl(sampleIdl, historyIdl *TreeNode) *TreeNode {
	if historyIdl == nil || sampleIdl == nil {
		return sampleIdl
	}
	//虽然当前root根节点 指针地址不一致 ，但是child指针地址是一致的 但是在这里就是需要修改child  所以最后还是通过序列化和反序列化完成吧
	mergeIdl := Deserializes(context.Background(), Serialize(context.Background(), sampleIdl))
	mergeIdl.Child[0].Child = mergeNodeArray(mergeIdl.Child[0].Child, historyIdl.Child[0].Child)
	mergeIdl.Child[1].Child = mergeNodeArray(mergeIdl.Child[1].Child, historyIdl.Child[1].Child)
	mergeIdl.Child[2].Child = mergeNodeArray(mergeIdl.Child[2].Child, historyIdl.Child[2].Child)
	return mergeIdl
}

// MergeNodeArray  将t1树合并到t当中
func mergeNodeArray(t, t1 []*TreeNode) []*TreeNode {
	if t == nil {
		return t1
	}
	if t1 == nil {
		return t
	}
	//判断节点是否相等 如果相等继续递归merge 不相等进行合并排序
	left, right := 0, 0
	for left < len(t) && right < len(t1) { //双指针走
		if *(t[left].Info) == *(t1[right].Info) {
			t[left].Child = mergeNodeArray(t[left].Child, t1[right].Child) //当两个节点一致 进行递归
			left++
			right++
		} else {
			left++
		}
	}
	if right < len(t1) {
		t = append(t, t1[right:]...)
		//对t 中key进行排序 进行排序
		t = SortSliceNodeInfo(t)
	}
	return t
}

func SortSliceNodeInfo(t []*TreeNode) []*TreeNode {

	m := make(map[string]*TreeNode) //key是node.Info.ParamKey

	list := make([]string, 0)
	for _, node := range t {
		m[node.Info.ParamKey+node.Info.ParamType] = node //参数key+type当做key 避免key一样 但是val不一致
	}

	for key, _ := range m { //val是 node.Info.ParamKey 进行排序
		list = append(list, key)
	}
	sort.Strings(list)

	res := make([]*TreeNode, 0)
	for _, key := range list {
		res = append(res, m[key])
	}
	return res
}

// 判断idl内容是否一致 如果不一致 返回false
func isSameIdl(mergeIdl, sampleIdl *TreeNode) bool {
	mergeSer := Serialize(context.Background(), mergeIdl)
	sampleSer := Serialize(context.Background(), sampleIdl)
	return mergeSer == sampleSer
}
