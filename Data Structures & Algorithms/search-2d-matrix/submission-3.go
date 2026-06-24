func searchMatrix(matrix [][]int, target int) bool {
	m,n := len(matrix), len(matrix[0])
	left, right := 0, m*n -1
	for left <= right {
		mid:= (left+right)/2
		value:= matrix[mid/n][mid%n]
		if value == target {
			return true
		} else if value>target {
			right = mid-1
		} else {
			left = mid+1
		}
	}
	return false
}
