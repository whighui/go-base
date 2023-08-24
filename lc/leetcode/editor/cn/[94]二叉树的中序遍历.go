//给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,3,2]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
//
// Related Topics 栈 树 深度优先搜索 二叉树 👍 1877 👎 0

package doc

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//func inorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//	result := make([]int, 0)
//	var inorder func(node *TreeNode)
//	inorder = func(node *TreeNode) {
//		if node == nil {
//			return
//		}
//		inorder(node.Left)
//		result = append(result, node.Val)
//		inorder(node.Right)
//	}
//	inorder(root)
//	return result
//}
// lc-94 使用二叉树迭代遍历  同样来使用前序遍历 类似于栈形式来完成呢
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1] //这里边本来就是嵌套到最后一层了 所以需要拿出来重新判断一下呗
		result = append(result, root.Val)
		root = root.Right
		stack = stack[:len(stack)-1]
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
