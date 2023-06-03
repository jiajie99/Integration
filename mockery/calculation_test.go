package mockery

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMockCalculation_Add(t *testing.T) {
	mockCalculation := NewMockCalculation(t)
	mockCalculation.EXPECT().Add(mock.Anything, mock.Anything).Return(6).Once()
	result := mockCalculation.Add(1, 2)
	assert.Equal(t, 6, result)
	// Another way.
	mockCalculation.On("Add", mock.Anything, mock.Anything).Return(7).Once()
	result = mockCalculation.Add(1, 2)
	assert.Equal(t, 7, result)
}
