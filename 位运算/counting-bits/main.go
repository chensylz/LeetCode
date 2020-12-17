package counting_bits

/**
338. 比特位计数
https://leetcode-cn.com/problems/counting-bits/
 */

func countBits(num int) []int {
	result := make([]int, num + 1)
	for i := 0; i < num; i++ {
		result[i] = count1(i)
	}
	return result
}


func count1(num int) int {
	result := 0
	for num != 0 {
		num = num & (num - 1)
		result += 1
	}
	return result
}