package first_bad_version

/**
278. 第一个错误的版本
https://leetcode-cn.com/problems/first-bad-version/
 */

func firstBadVersion(n int) int {
	start := 0
	end := n
	for start + 1 < end {
		mid := start + (end - start) / 2
		if isBadVersion(mid) {
			end = mid
		} else if isBadVersion(mid) == false {
			start = mid
		}
	}

	if isBadVersion(start) {
		return start
	}
	return end
}