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
