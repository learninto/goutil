package excelize

import (
	"testing"

	"github.com/learninto/goutil/test/assert"
)

func TestCoordinatesToCellName(t *testing.T) {
	s, e := CoordinatesToCellName(1, 1)
	assert.Nil(t, e)
	assert.Equal(t, s, "A1")
}
