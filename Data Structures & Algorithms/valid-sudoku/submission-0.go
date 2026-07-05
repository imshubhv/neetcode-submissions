func isValidSudoku(board [][]byte) bool {
    row:= 0
    for row < 9 {
        numbers:= map[byte]byte{}
        for i:=0; i<9; i++ {
            val:= board[row][i]
            if numbers[val] == val{
                return false
            } else if val != '.'{
                numbers[val] = val
            }
        }
        row++
    }

    column:= 0
    for column < 9 {
        numbers:= map[byte]byte{}
        for i:=0; i<9; i++ {
            val:= board[i][column]
            if numbers[val] == val{
                return false
            } else if val != '.'{
                numbers[val] = val
            }
            fmt.Println(numbers)
        }
        column++
    }
    for boxRow:= 0; boxRow<9; boxRow+=3 {
        for boxCol:= 0; boxCol<9; boxCol+=3 {
            numbers:= map[byte]byte{}
            for i:= boxRow ; i< boxRow+3 ; i++ {
                for j:=boxCol; j< boxCol+3 ; j++ {
                    val:= board[i][j]
                    if numbers[val] == val{
                        return false
                    } else if val != '.'{
                        numbers[val] = val
                    }
                }
            }
        }
    }
    return true
}
