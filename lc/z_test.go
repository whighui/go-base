package lc

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	permute([]int{1, 2, 3})
}

func TestMinLength(t *testing.T) {
	nums := make([]int, 0)
	nums = append(nums, 1, 2, 3, 4, 5, 5, 5, 1)
	fmt.Println(minLength(nums))
}

func TestRemoveKdigits(t *testing.T) {
	fmt.Println(removeKdigits("12532", 2))
}

// [[2],[2],[2,6],[1],[1,5],[1,2],[1],[2]]
func TestLru(t *testing.T) {
	lRUCache := Constructor(2)
	lRUCache.Get(2)    //-1
	lRUCache.Put(2, 6) //null
	lRUCache.Get(1)    //-1
	lRUCache.Put(1, 5) //null
	lRUCache.Put(1, 2) //null
	lRUCache.Get(1)    //2
	lRUCache.Get(2)    //6
	//[null,-1,null,-1,null,null,2,6]

}

func TestFindUnsortedSubarray(t *testing.T) {
	nums := []int{1, 3, 2, 2, 2}
	fmt.Println(findUnsortedSubarray(nums))
}

func TestMaxSubArray1(t *testing.T) {
	nums := []int{-1, -2, -3, 4}
	fmt.Println(maxSubArray(nums))
}

func TestLengthOfLIS(t *testing.T) {
	maxLength := lengthOfLIS([]int{1, 2, 3, 5, 4})
	fmt.Println(maxLength)
}

// 最长递增子序列的个数
func TestLengthOfLISCount(t *testing.T) {
	maxLengthCount := lengthOfLISCount([]int{1, 1, 1, 1})

	fmt.Println(maxLengthCount)
}

func TestQueue(t *testing.T) {
	queue := make([]*TreeNode, 0)
	queue = append(queue, &TreeNode{Val: 1}, nil)
	fmt.Println(len(queue))
	fmt.Println(queue)
}

func TestIsSymmetric(t *testing.T) {
	root := &TreeNode{2, &TreeNode{
		Val:   3,
		Right: &TreeNode{Val: 4},
	}, &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 4},
	}}

	fmt.Println(isSymmetric(root))
}
