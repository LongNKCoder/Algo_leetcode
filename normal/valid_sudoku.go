package main

import (
	"fmt"
	"strconv"
)

//Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
//
//Each row must contain the digits 1-9 without repetition.
//Each column must contain the digits 1-9 without repetition.
//Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
//Note:
//
//A Sudoku board (partially filled) could be valid but is not necessarily solvable.
//Only the filled cells need to be validated according to the mentioned rules.

func main() {
	// testcase
	fmt.Println(isValidSudokuEnhance([][]string{
		[]string{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		[]string{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		[]string{".", "9", "8", ".", ".", ".", ".", "6", "."},
		[]string{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		[]string{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		[]string{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		[]string{".", "6", ".", ".", ".", ".", "2", "8", "."},
		[]string{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		[]string{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}))
	fmt.Println(isValidSudokuEnhance([][]string{
		[]string{"8", "3", ".", ".", "7", ".", ".", ".", "."},
		[]string{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		[]string{".", "9", "7", ".", ".", ".", ".", "6", "."},
		[]string{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		[]string{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		[]string{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		[]string{".", "6", ".", ".", ".", ".", "2", "8", "."},
		[]string{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		[]string{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	})) // dup on col
	fmt.Println(isValidSudokuEnhance([][]string{
		[]string{"1", "3", ".", ".", "7", ".", ".", ".", "."},
		[]string{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		[]string{"8", "9", "8", ".", ".", ".", ".", "6", "."},
		[]string{"2", ".", ".", ".", "6", ".", ".", ".", "3"},
		[]string{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		[]string{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		[]string{".", "6", ".", ".", ".", ".", "2", "8", "."},
		[]string{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		[]string{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	})) // dup on row
	fmt.Println(isValidSudokuEnhance([][]string{
		[]string{".", "2", ".", ".", ".", "1", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", ".", "5", ".", "."},
		[]string{".", ".", ".", "1", ".", ".", "8", ".", "3"},
		[]string{"4", ".", "9", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", "2", ".", ".", ".", ".", ".", "."},
		[]string{"7", ".", "6", ".", ".", ".", ".", ".", "."},
		[]string{"9", ".", ".", ".", ".", "4", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", ".", ".", ".", "."},
	})) // dup on box
	fmt.Println(isValidSudokuEnhance([][]string{
		[]string{".", ".", ".", ".", "5", ".", ".", "1", "."},
		[]string{".", "4", ".", "3", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", "3", ".", ".", "1"},
		[]string{"8", ".", ".", ".", ".", ".", ".", "2", "."},
		[]string{".", ".", "2", ".", "7", ".", ".", ".", "."},
		[]string{".", "1", "5", ".", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", "2", ".", ".", "."},
		[]string{".", "2", ".", "9", ".", ".", ".", "6", "."},
		[]string{".", ".", "3", ".", ".", ".", ".", ".", "."},
	})) // dup on box
}

func isValidSudoku(board [][]string) bool {
	var (
		mapRows = make(map[int][]string)
		mapCols = make(map[int][]string)
		mapMatr = make(map[int][]string)
	)
	// init map
	for i := 0; i < len(board); i++ {
		// init row
		mapRows[i+1] = board[i]
		for j := 0; j < len(board[i]); j++ {
			// init column
			if _, ok := mapCols[i+1]; !ok {
				mapCols[j+1] = make([]string, 0)
			}
			mapCols[j+1] = append(mapCols[j+1], board[i][j])
			// init matrix
			switch {
			case i >= 0 && i < 3 && j < 3:
				mapMatr[1] = append(mapMatr[1], board[i][j])
			case i >= 3 && i < 6 && j < 3:
				mapMatr[2] = append(mapMatr[2], board[i][j])
			case i >= 6 && j < 3:
				mapMatr[3] = append(mapMatr[3], board[i][j])
			case i >= 0 && i < 3 && j < 6:
				mapMatr[4] = append(mapMatr[4], board[i][j])
			case i >= 3 && i < 6 && j < 6:
				mapMatr[5] = append(mapMatr[5], board[i][j])
			case i >= 6 && j < 6:
				mapMatr[6] = append(mapMatr[6], board[i][j])
			case i >= 0 && i < 3 && j < 9:
				mapMatr[7] = append(mapMatr[7], board[i][j])
			case i >= 3 && i < 6 && j < 9:
				mapMatr[8] = append(mapMatr[8], board[i][j])
			case i >= 6 && j < 8:
				mapMatr[9] = append(mapMatr[9], board[i][j])
			}
		}
	}
	// check row
	for _, row := range mapRows {
		if !checkValid(row) {
			return false
		}
	}
	// check column
	for _, col := range mapCols {
		if !checkValid(col) {
			return false
		}
	}
	// check matrix
	for _, mat := range mapMatr {
		if !checkValid(mat) {
			return false
		}
	}
	return true
}

func checkValid(row []string) bool {
	mapExistNum := make(map[int]bool)
	for _, num := range row {
		intNum, err := strconv.Atoi(num)
		if err != nil {
			continue
		}
		if _, ok := mapExistNum[intNum]; ok {
			return false
		} else {
			mapExistNum[intNum] = true
		}
	}
	return true
}

// Time complexity: O(n^2)
// Space complexity: O(n^2)

// solve problem with better solution
func isValidSudokuEnhance(board [][]string) bool {
	var (
		rows = make([]map[string]bool, 9)
		cols = make([]map[string]bool, 9)
		mats = make([]map[string]bool, 9)
	)
	for i := 0; i < 9; i++ {
		rows[i] = make(map[string]bool)
		cols[i] = make(map[string]bool)
		mats[i] = make(map[string]bool)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num == "." {
				continue
			}
			if _, ok := rows[i][num]; ok {
				return false
			}
			if _, ok := cols[j][num]; ok {
				return false
			}
			matIndex := (i/3)*3 + j/3
			if _, ok := mats[matIndex][num]; ok {
				return false
			}
			rows[i][num] = true
			cols[j][num] = true
			mats[matIndex][num] = true
		}
	}
	return true
}

/*
example:

"5", "3", ".", ".", "7", ".", ".", ".", "."
"6", ".", ".", "1", "9", "5", ".", ".", "."
".", "9", "8", ".", ".", ".", ".", "6", "."
"8", ".", ".", ".", "6", ".", ".", ".", "3"
"4", ".", ".", "8", ".", "3", ".", ".", "1"
"7", ".", ".", ".", "2", ".", ".", ".", "6"
".", "6", ".", ".", ".", ".", "2", "8", "."
".", ".", ".", "4", "1", "9", ".", ".", "5"
".", ".", ".", ".", "8", ".", ".", "7", "9"

"8", "3", ".", ".", "7", ".", ".", ".", "."
"6", ".", ".", "1", "9", "5", ".", ".", "."
".", "9", "7", ".", ".", ".", ".", "6", "."
"8", ".", ".", ".", "6", ".", ".", ".", "3"
"4", ".", ".", "8", ".", "3", ".", ".", "1"
"7", ".", ".", ".", "2", ".", ".", ".", "6"
".", "6", ".", ".", ".", ".", "2", "8", "."
".", ".", ".", "4", "1", "9", ".", ".", "5"
".", ".", ".", ".", "8", ".", ".", "7", "9"

"1", "3", ".", ".", "7", ".", ".", ".", "."
"6", ".", ".", "1", "9", "5", ".", ".", "."
"8", "9", "8", ".", ".", ".", ".", "6", "."
"2", ".", ".", ".", "6", ".", ".", ".", "3"
"4", ".", ".", "8", ".", "3", ".", ".", "1"
"7", ".", ".", ".", "2", ".", ".", ".", "6"
".", "6", ".", ".", ".", ".", "2", "8", "."
".", ".", ".", "4", "1", "9", ".", ".", "5"
".", ".", ".", ".", "8", ".", ".", "7", "9"

".", "2", ".", ".", ".", "1", ".", ".", "."
".", ".", ".", ".", ".", ".", "5", ".", "."
".", ".", ".", "1", ".", ".", "8", ".", "3"
"4", ".", "9", ".", ".", ".", ".", ".", "."
".", ".", ".", ".", ".", ".", ".", ".", "."
".", ".", "2", ".", ".", ".", ".", ".", "."
"7", ".", "6", ".", ".", ".", ".", ".", "."
"9", ".", ".", ".", ".", "4", ".", ".", "."
".", ".", ".", ".", ".", ".", ".", ".", "."

[[".", ".", ".", ".", "5", ".", ".", "1", "."],
[".", "4", ".", "3", ".", ".", ".", ".", "."],
[".", ".", ".", ".", ".", "3", ".", ".", "1"],
["8", ".", ".", ".", ".", ".", ".", "2", "."],
[".", ".", "2", ".", "7", ".", ".", ".", "."],
[".", "1", "5", ".", ".", ".", ".", ".", "."],
[".", ".", ".", ".", ".", "2", ".", ".", "."],
[".", "2", ".", "9", ".", ".", ".", "6", "."],
[".", ".", "3", ".", ".", ".", ".", ".", "."]]
*/
