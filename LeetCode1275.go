package main

func tictactoe(moves [][]int) string {
	var board [3][3]int
	for i:=0;i<3;i++ {
		for j:=0;j<3;j++ {
			board[i][j] = -1
		}
	}


	var numMoves = len(moves)
	for i:= 0;i<numMoves;i++ {
		var row = moves[i][0]
		var col = moves[i][1]
		if i%2 == 0 {
			board[row][col] = 0
		} else {
			board[row][col] = 1
		}
	}

	return getResult(board, numMoves)
}

func getResult(board [3][3]int, numMoves int) string {
	var result string
	for i:=0;i<3;i++ {
		result = checkRow(board, i)
		if result != "" {
			return result
		}
		result = checkCol(board, i)
		if result != "" {
			return result
		}
	}

	result = checkDiags(board)
	if result != "" {
		return result
	}

	if numMoves == 9 {
		return "Draw"
	}

	return "Pending"
}

func checkRow(board [3][3]int, rowNum int) string {
	if board[rowNum][0] == board[rowNum][1] && board[rowNum][1] == board[rowNum][2] {
		if board[rowNum][0] == 0 {
			return "A"
		} else if board[rowNum][0] == 1 {
			return "B"
		}
	}

	return ""
}

func checkCol(board [3][3]int, colNum int) string {
	if board[0][colNum] == board[1][colNum] && board[1][colNum] == board[2][colNum] {
		if board[0][colNum] == 0 {
			return "A"
		} else if board[0][colNum] == 1 {
			return "B"
		}
	}

	return ""
}

func checkDiags(board [3][3]int) string {
	if (board[0][0] == board[1][1] && board[1][1] == board[2][2]) ||
		(board[2][0] == board[1][1] && board[1][1] == board[0][2]){
		if board[1][1] == 0 {
			return "A"
		} else if board[1][1] == 1 {
			return "B"
		}
	}

	return ""
}
