package excelize

import (
	"io"

	excelize "github.com/xuri/excelize/v2"
)

// NewFile 创建文件
func NewFile() *excelize.File {
	return excelize.NewFile()
}

// OpenFile 打开文件
func OpenFile(filename string) (*excelize.File, error) {
	return excelize.OpenFile(filename)
}

// OpenReader 打开文件流
func OpenReader(r io.Reader) (*excelize.File, error) {
	return excelize.OpenReader(r)
}

// CoordinatesToCellName converts [X, Y] coordinates to alpha-numeric cell
// name or returns an error.
//
// Example:
//
//    CoordinatesToCellName(1, 1) // returns "A1", nil
//
func CoordinatesToCellName(col, row int) (string, error) {
	return excelize.CoordinatesToCellName(col, row)
}
