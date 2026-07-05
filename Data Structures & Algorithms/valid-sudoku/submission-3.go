func isValidSudoku(board [][]byte) bool {
    var rows [9]int
    var cols [9]int
    var boxes [9]int

    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {

            if board[r][c] == '.' {
                continue
            }

            digit := board[r][c] - '1'
            mask := 1 << digit
            box := (r/3)*3 + c/3

            if rows[r]&mask != 0 ||
                cols[c]&mask != 0 ||
                boxes[box]&mask != 0 {
                return false
            }

            rows[r] |= mask
            cols[c] |= mask
            boxes[box] |= mask
        }
    }

    return true
}
