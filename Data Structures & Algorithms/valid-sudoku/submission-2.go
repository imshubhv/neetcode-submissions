func isValidSudoku(board [][]byte) bool {
    row:= 0
    for row < 9 {
        numbers:= map[byte]bool{}
        for i:=0; i<9; i++ {
            if board[row][i] == '.' {
                continue
            }
            val:= board[row][i]
            if numbers[val]{
                return false
            } else if val != '.'{
                numbers[val] = true
            }
        }
        row++
    }

    column:= 0
    for column < 9 {
        numbers:= map[byte]bool{}
        for i:=0; i<9; i++ {
            if board[i][column] == '.' {
                continue
            }
            val:= board[i][column]
            if numbers[val]{
                return false
            } else if val != '.'{
                numbers[val] = true
            }
        }
        column++
    }
    for boxRow:= 0; boxRow<9; boxRow+=3 {
        for boxCol:= 0; boxCol<9; boxCol+=3 {
            numbers:= map[byte]bool{}
            for i:= boxRow ; i< boxRow+3 ; i++ {
                for j:=boxCol; j< boxCol+3 ; j++ {
                    if board[i][j] == '.' {
                        continue
                    }
                    val:= board[i][j]
                    if numbers[val]{
                        return false
                    } else if val != '.'{
                        numbers[val] = true
                    }
                }
            }
        }
    }
    return true
}
