package gomock

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMockCalculation_Add(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCalculation := NewMockCalculation(ctrl)
	mockCalculation.EXPECT().Add(gomock.Any(), gomock.Any()).Return(6).Times(1)
	result := mockCalculation.Add(1, 2)
	assert.Equal(t, 6, result)
}
