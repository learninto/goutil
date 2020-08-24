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
