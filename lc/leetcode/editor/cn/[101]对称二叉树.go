//给你一个二叉树的根节点 root ， 检查它是否轴对称。
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,2,3,4,4,3]
//输出：true
//
//
// 示例 2：
//
//
//输入：root = [1,2,2,null,3,null,3]
//输出：false
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [1, 1000] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：你可以运用递归和迭代两种方法解决这个问题吗？
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 2542 👎 0

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
//给定二叉树 判断是否对称
//递归法和迭代法能来完成吗
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}

// 需要进行确定思路呗 查看树的形状呢 因为在这里就是用一个函数好像是完成不了呢
func compare(leftNode, rightNode *TreeNode) bool {
	if leftNode == nil && rightNode == nil {
		return true
	}
	if leftNode == nil && rightNode != nil {
		return false
	}
	if rightNode == nil && leftNode != nil {
		return false
	}
	if leftNode.Val != rightNode.Val {
		return false
	}
	return compare(leftNode.Left, rightNode.Right) && compare(leftNode.Right, rightNode.Left)
}

// 使用BFS层序遍历好像也是可以完成此题的呢
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root.Left, root.Right)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size/2; i++ {
			leftNode, rightNode := queue[i], queue[size-1-i]
			if leftNode == nil && rightNode == nil {
				continue
			}
			if leftNode == nil || rightNode == nil || leftNode.Val != rightNode.Val {
				return false
			}
			queue = append(queue, leftNode.Left, leftNode.Right, rightNode.Left, rightNode.Right)
		}
		queue = queue[size:]
		size = len(queue)
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
