package gomock

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMockCalculation_Add(t *testing.T) {
	t.Run("mock function", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockCalculation := NewMockcalculation(ctrl)
		mockCalculation.EXPECT().add(gomock.Any(), gomock.Any()).Return(6).Times(1)
		result := mockCalculation.add(1, 2)
		assert.Equal(t, 6, result)
	})
	t.Run("mock function (uncovered closure)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockCalculation := NewMockcalculation(ctrl)
		mockCalculation.EXPECT().addFunc(gomock.Any(), gomock.Any(), gomock.Any()).Return(4)
		e := &example{
			c: mockCalculation,
		}
		result := e.executeAdd(1, 1)
		assert.Equal(t, 4, result)
	})
	t.Run("mock function (covered closure)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockCalculation := NewMockcalculation(ctrl)
		f := func(a, b int, f func(a, b int) int) int {
			return f(a, b)
		}
		mockCalculation.EXPECT().addFunc(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(f)
		e := &example{
			c: mockCalculation,
		}
		result := e.executeAdd(1, 1)
		assert.Equal(t, 2, result)
	})
}
