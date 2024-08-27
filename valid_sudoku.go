func getBoxIndex(row int, col int) int {
    return ((row / 3) * 3) + (col / 3)
}


func isValidSudoku(board [][]byte) bool {
    subBoxHashes := make([]map[byte]bool, 9)
    columnHashes := make([]map[byte]bool, 9)
    rowHashes := make([]map[byte]bool, 9)
    var subBoxIndex int
    var currentValue byte

    for i := range subBoxHashes {
        columnHashes[i] = make(map[byte]bool)
        rowHashes[i] = make(map[byte]bool)
        subBoxHashes[i] = make(map[byte]bool)
    } 

    for rowIndex, row := range board {
        for colIndex, _ := range row {

            subBoxIndex = getBoxIndex(rowIndex, colIndex)
            currentValue = board[rowIndex][colIndex]

            if currentValue == '.' {
                continue
            }

            if columnHashes[rowIndex][currentValue] ||
                rowHashes[colIndex][currentValue] ||
                subBoxHashes[subBoxIndex][currentValue] {
                return false
            }

            columnHashes[rowIndex][currentValue] = true
            rowHashes[colIndex][currentValue] = true
            subBoxHashes[subBoxIndex][currentValue] = true
            
        }
    }

    return true

}
