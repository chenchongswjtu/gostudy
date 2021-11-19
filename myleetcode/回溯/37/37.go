package main

import "fmt"

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	solveSudoku(board)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if j == 0 {
				fmt.Print(board[i][j] - '0')
			} else {
				fmt.Print("  ", board[i][j]-'0')
			}
		}
		fmt.Println()
	}
}

type position struct {
	x int
	y int
}

func solveSudoku(board [][]byte) {
	var pos []position
	var find bool

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				pos = append(pos, position{x: i, y: j})
			}
		}
	}

	putSudoku(&board, pos, 0, &find)
}

// 回溯
func putSudoku(board *[][]byte, pos []position, index int, success *bool) {
	if *success {
		return
	}

	if index == len(pos) {
		*success = true
		return
	}

	for i := 1; i < 10; i++ { // 从1到9进行尝试
		if checkSudoku(board, pos[index], i) && !*success {
			(*board)[pos[index].x][pos[index].y] = byte(i) + '0'
			putSudoku(board, pos, index+1, success)
			if *success {
				return
			}
			(*board)[pos[index].x][pos[index].y] = '.'
		}
	}
}

func checkSudoku(board *[][]byte, pos position, val int) bool {
	// 判断横行是否有重复数字
	for i := 0; i < len((*board)[0]); i++ {
		if (*board)[pos.x][i] != '.' && int((*board)[pos.x][i]-'0') == val {
			return false
		}
	}

	// 判断竖行是否有重复数字
	for i := 0; i < len(*board); i++ {
		if (*board)[i][pos.y] != '.' && int((*board)[i][pos.y]-'0') == val {
			return false
		}
	}

	// 判断九宫格是否有重复数字
	startX, startY := pos.x-pos.x%3, pos.y-pos.y%3 // 九宫格的左上角坐标
	for i := startX; i < startX+3; i++ {
		for j := startY; j < startY+3; j++ {
			if (*board)[i][j] != '.' && int((*board)[i][j]-'0') == val {
				return false
			}
		}
	}

	return true
}
