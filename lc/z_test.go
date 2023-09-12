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

func TestMaxSubArray1(t *testing.T) {
	nums := []int{-1, -2, -3, 4}
	fmt.Println(maxSubArray(nums))
}

func TestLengthOfLIS(t *testing.T) {
	maxLength := lengthOfLIS([]int{1, 2, 3, 5, 4})
	fmt.Println(maxLength)
}

//最长递增子序列的个数
func TestLengthOfLISCount(t *testing.T) {
	maxLengthCount := lengthOfLISCount([]int{1, 1, 1, 1})

	fmt.Println(maxLengthCount)
}
