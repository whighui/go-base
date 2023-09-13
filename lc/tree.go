package lc

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
	二叉树 前序：根左右
		  中序：左根右
          后续：左右根
          已知前序和中序可以确定一颗二叉树
		  已知中序和后序可以确定一颗二叉树

*/

// lc-144 二叉树前序遍历 递归完成
func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return result
}

// lc-144 二叉树前序遍历 迭代来完成 这里边就是需要使用栈来操作
func preorderTraversalLoop(root *TreeNode) []int {
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

		}
	}
	return result
}

// lc-94  二叉树中序遍历 使用递归来完成  左根右
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		result = append(result, node.Val)
		inorder(node.Right)
	}
	return result
}

// lc-94 使用二叉树迭代遍历  同样来使用前序遍历 类似于栈形式来完成呢
func inorderTraversalLoop(root *TreeNode) []int {
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
		result = append(result, root.Val)
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return result
}

// lc-100 想同的树
func isSameTree(p *TreeNode, q *TreeNode) bool {

	return false
}

// lc-113  路径总和II
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	var dfs func(node *TreeNode, list []int, sum int)

	dfs = func(node *TreeNode, list []int, sum int) {
		if node == nil {
			return
		}
		sum = sum + node.Val
		list = append(list, node.Val)
		if node.Left == nil && node.Right == nil && sum == 0 {
			temp := make([]int, len(list))
			copy(temp, list)
			result = append(result, temp)
			return
		}
		dfs(root.Left, list, sum)
		dfs(root.Right, list, sum)
		list = list[0 : len(list)-1]
		sum = sum - node.Val
	}
	list := make([]int, 0)
	dfs(root, list, targetSum)
	return result
}

// 合并二叉树
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	var val1, val2 int
	if root1 != nil {
		val1 = root1.Val
	}
	if root2 != nil {
		val2 = root2.Val
	}
	root := &TreeNode{Val: val1 + val2}
	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)
	return root
}
