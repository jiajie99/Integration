package validator

import (
	"encoding/json"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

// 传了类型的非零值，所以能过 required 校验
var case1 = `{
	"field1": "a",
	"field2": ""
}`

// 显式传了类型零值，命中 omitempty，不再进行后续校验
var case2 = `{
	"field1": 0,
	"field2": null
}`

var case3 = `{
	"field1": [
		{
			"field1": "a",
			"field2": ""
		}
	]
}`

type Struct1 struct {
	Field1 string  `json:"field1" validate:"required"`
	Field2 *string `json:"field2" validate:"required"`
}

type Struct2 struct {
	Field1 int  `json:"field1" validate:"omitempty,oneof=1 2 3"`
	Field2 *int `json:"field2" validate:"omitempty,min=1"`
}

type Struct3 struct {
	Field1 []Struct1 `json:"field1" validate:"required,dive"`
}

func TestRequired(t *testing.T) {
	validate := validator.New()
	t.Run("test required", func(t *testing.T) {
		struct1 := Struct1{}
		err := json.Unmarshal([]byte(case1), &struct1)
		assert.NoError(t, err)
		err = validate.Struct(struct1)
		assert.NoError(t, err)
	})

	t.Run("test required and omitempty", func(t *testing.T) {
		struct2 := Struct2{}
		err := json.Unmarshal([]byte(case2), &struct2)
		assert.NoError(t, err)
		err = validate.Struct(struct2)
		assert.NoError(t, err)
	})

	t.Run("test dive", func(t *testing.T) {
		struct3 := Struct3{}
		err := json.Unmarshal([]byte(case3), &struct3)
		assert.NoError(t, err)
		err = validate.Struct(struct3)
		assert.NoError(t, err)
	})
}
