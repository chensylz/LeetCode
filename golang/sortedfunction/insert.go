package sortedfunction

// Insert 插入排序.
// 这个想法是将数组分成两个子集——排序子集和未排序子集。
// 最初，排序子集仅包含索引处的一个第一个元素 0.然后对于每次迭代，插入排序从未排序的子集中删除下一个元素，
// 在已排序的子集中找到它所属的位置并将其插入那里。它重复直到没有输入元素
func Insert(nums []int) {
	for i := 1; i < len(nums); i++ {
		current := nums[i]
		var index int
		for index = i - 1; index >= 0; index-- {
			if nums[index] > current {
				nums[index+1] = nums[index]
			} else {
				break
			}
		}
		nums[index+1] = current
	}
}
