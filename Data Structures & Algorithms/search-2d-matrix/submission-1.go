func searchMatrix(matrix [][]int, target int) bool {
	row:= len(matrix)
	column:= len(matrix[0])
	if target > matrix[row-1][column-1] || target < matrix[0][0]{
		return false
	}
	for i:= 0; i<row; i++{
		if target >= matrix[i][0] && target <= matrix[i][column-1] {
			left, right:= 0, column-1
			// fmt.Println(matrix[left][right])
			for left<=right {
				mid:= (right+left)/2
				// fmt.Println(matrix[i][mid])
				if matrix[i][mid] == target {
					return true
				} else if matrix[i][mid] > target {
					right = mid-1
				}else {
					left = mid+1
				}
			}
		}
	}
	// fmt.Println("There")
	return false
}
