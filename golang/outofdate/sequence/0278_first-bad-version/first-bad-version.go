package _278_first_bad_version


/**
题目:
你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。
假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
你可以通过调用bool isBadVersion(version)接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。

example1:
	输入：n = 5, bad = 4
	输出：4
	解释：
	调用 isBadVersion(3) -> false
	调用 isBadVersion(5)-> true
	调用 isBadVersion(4)-> true
	所以，4 是第一个错误的版本。

example2:
	输入：n = 1, bad = 1
	输出：1

思路:
	仍然是二分查找，不过是找到相同target的第一个
 */

func firstBadVersion(n int) int {
	return binarySearch(n)
}

func binarySearch(n int) int {
	start, end := 0, n
	for start < end {
		mid := (end - start) / 2 + start
		if isBadVersion(mid) {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return end
}


func isBadVersion(n int) bool {
	return true
}