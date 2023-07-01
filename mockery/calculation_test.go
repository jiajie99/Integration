package mockery

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMockCalculation_Add(t *testing.T) {
	mockCal := newMockCalculation(t)
	mockCal.EXPECT().Add(mock.Anything, mock.Anything).Return(6).Once()
	result := mockCal.Add(1, 2)
	assert.Equal(t, 6, result)
	// Another way.
	mockCal.On("Add", mock.Anything, mock.Anything).Return(7).Once()
	result = mockCal.Add(1, 2)
	assert.Equal(t, 7, result)
}
