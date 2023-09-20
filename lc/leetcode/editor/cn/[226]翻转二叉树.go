//给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [4,2,7,1,3,6,9]
//输出：[4,7,2,9,6,3,1]
//
//
// 示例 2：
//
//
//
//
//输入：root = [2,1,3]
//输出：[2,3,1]
//
//
// 示例 3：
//
//
//输入：root = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点数目范围在 [0, 100] 内
// -100 <= Node.val <= 100
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1682 👎 0

package doc

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
// */

//翻转二叉树 三种方式来进行遍历
//递归 、 dfs 、bfs
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left //父
	invertTree(root.Left)                         //左
	invertTree(root.Right)                        //右
	return root
}

// 深度优先遍历 DFS 利用栈来实现
func invertTreeDfs(root *TreeNode) *TreeNode {
	//依然采用前序遍历方式来完成迭代
	stack := make([]*TreeNode, 0)
	result := root //因为这里边root就是一直在改变 所以需要进行最开是进行保存结果
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root.Left, root.Right = root.Right, root.Left
			root = root.Left
		}
		root = stack[len(stack)-1]   //栈顶元素
		stack = stack[:len(stack)-1] //弹栈
		root = root.Right            //去寻找右孩子
	}
	return result
}

// invertTree 广度优先遍历 使用队列来完成
func invertTreeBfs(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
		size = len(queue)
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
