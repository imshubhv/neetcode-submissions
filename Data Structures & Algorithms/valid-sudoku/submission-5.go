func isValidSudoku(board [][]byte) bool {
    var rows [9]int
    var col [9]int
    var box [9]int

    for i:=0 ; i<9; i++ {
        for j:=0; j<9; j++ {
            val := board[i][j]
            if val == '.'{
                continue
            }
            val -= '1'
            mask:= 1 << val
            k:= (i/3)*3+(j/3)
            if rows[i] & mask !=0 || col[j] & mask !=0 || box[k] & mask !=0 {
                return false
            }
            rows[i] = rows[i] | mask
            col[j] = col[j] | mask
            box[k] = box[k] | mask
        }
    }
    return true
}