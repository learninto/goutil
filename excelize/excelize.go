package excelize

import (
	"io"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// NewFile
func NewFile() *excelize.File {
	return excelize.NewFile()
}

// OpenFile
func OpenFile(filename string) (*excelize.File, error) {
	return excelize.OpenFile(filename)
}

// OpenReader
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
