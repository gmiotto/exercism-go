package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix 1
type Matrix interface {
	Rows() [][]int
	Cols() [][]int
	Set(row, col, val int) bool
}

// typeMatrix 1
type typeMatrix [][]int

// New 1
func New(input string) (Matrix, error) {
	m := typeMatrix{}
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		arrRowTemp := strings.Split(row, " ")
		arrRow := []int{}
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.Atoi(arrRowItem)
			if err != nil {
				return nil, errors.New("Invalid input")
			}
			arrRow = append(arrRow, arrItemTemp)
		}
		m = append(m, arrRow)
	}
	if !verifyMatrixIsEven(m) {
		return nil, errors.New("Invalid matrix size")
	}
	return m, nil
}

// Rows 1
func (m typeMatrix) Rows() [][]int {
	rows := make([][]int, len(m))
	for i, row := range m {
		for _, item := range row {
			rows[i] = append(rows[i], item)
		}
	}
	return rows
}

// Cols 1
func (m typeMatrix) Cols() [][]int {
	cols := make([][]int, len(m[0]))
	for _, row := range m {
		for j, item := range row {
			cols[j] = append(cols[j], item)
		}
	}
	return cols
}

// Set 1
func (m typeMatrix) Set(row, col, val int) bool {
	if len(m) > row && len(m[0]) > col && col >= 0 && row >= 0 {
		m[row][col] = val
		return true
	}
	return false
}

func verifyMatrixIsEven(m typeMatrix) bool {
	size := -1
	for _, row := range m {
		if size == -1 {
			size = len(row)
			continue
		}
		if size != len(row) {
			return false
		}
	}
	return true
}
