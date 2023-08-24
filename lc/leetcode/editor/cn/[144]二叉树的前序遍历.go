//给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,2,3]
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
// 示例 4：
//
//
//输入：root = [1,2]
//输出：[1,2]
//
//
// 示例 5：
//
//
//输入：root = [1,null,2]
//输出：[1,2]
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
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
//
// Related Topics 栈 树 深度优先搜索 二叉树 👍 1116 👎 0

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
// lc-144 二叉树前序遍历
//func preorderTraversal(root *TreeNode) []int {
//	result := make([]int, 0)
//	var preorder func(*TreeNode)
//	preorder = func(node *TreeNode) {
//		if node == nil {
//			return
//		}
//		result = append(result, node.Val)
//		preorder(node.Left)
//		preorder(node.Right)
//	}
//	preorder(root)
//	return result
//}

func preorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			result = append(result, root.Val)
			root = root.Left
		}
		//上述就是压入根节点和左孩子了 需要弹出找右孩子呢
		if len(stack) > 0 {
			root = stack[len(stack)-1].Right
			//	这里边需要重置一下stack长度
			stack = stack[:len(stack)-1]
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
