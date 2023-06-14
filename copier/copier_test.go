package copier

import (
	"testing"

	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
)

type Struct1 struct {
	Field1       string
	Field2       int
	DoubleField3 int
}

func (s1 *Struct1) DoubleField2() int {
	return s1.Field2 * 2
}

type Struct2 struct {
	Field1       string
	DoubleField2 int
	Field3       int
}

func (s2 *Struct2) DoubleField3(doubleField3 int) {
	s2.Field3 = doubleField3 * 2
}

func TestCopier(t *testing.T) {
	t.Run("copy from field to field with same name", func(t *testing.T) {
		struct1 := Struct1{
			Field1: "field1",
		}
		struct2 := Struct2{}
		copier.Copy(&struct2, &struct1)
		assert.Equal(t, "field1", struct2.Field1)
	})

	t.Run("copy from method to field with same name", func(t *testing.T) {
		struct1 := Struct1{
			DoubleField3: 10,
		}
		struct2 := Struct2{}
		copier.Copy(&struct2, &struct1)
		assert.Equal(t, 20, struct2.Field3)
	})

	t.Run("copy from field to method with same name", func(t *testing.T) {
		struct1 := Struct1{
			Field2: 10,
		}
		struct2 := Struct2{}
		copier.Copy(&struct2, &struct1)
		assert.Equal(t, 20, struct2.DoubleField2)
	})
}
