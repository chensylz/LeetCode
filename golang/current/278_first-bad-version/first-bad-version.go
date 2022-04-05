package _78_first_bad_version

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	return binarySearch(n)
}

func binarySearch(n int) int {
	start, end := 0, n
	for start < end {
		mid := (end-start)/2 + start
		if isBadVersion(mid) {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return end
}
