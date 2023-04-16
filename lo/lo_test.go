package lo

import (
	"strconv"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestLo(t *testing.T) {
	names := lo.Uniq([]string{"Samuel", "John", "Samuel"})
	assert.Equal(t, names, []string{"Samuel", "John"})

	even := lo.Filter([]int{1, 2, 3, 4}, func(x int, index int) bool {
		return x%2 == 0
	})
	assert.Equal(t, even, []int{2, 4})

	res := lo.Map([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})
	assert.Equal(t, res, []string{"1", "2", "3", "4"})
}
