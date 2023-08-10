package doc

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {

	//这里边就是返回数组下标呗
	m := make(map[int]int)

	for index, val := range nums {
		if _, ok := m[target-val]; ok {
			return []int{m[target-val], index}
		}
		m[val] = index
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
