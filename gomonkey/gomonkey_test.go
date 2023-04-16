package gomonkey

import (
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

type example struct {
	A, B int
}

func (e example) add() int {
	return e.A + e.B
}

func executeAdd(a, b int) int {
	e := example{A: a, B: b}
	return e.add()
}

// export GOARCH=amd64;go test -gcflags=all=-l
func TestGoMonkey(t *testing.T) {
	var e example
	patches := gomonkey.ApplyPrivateMethod(
		reflect.TypeOf(e), "add", func(e example) int {
			return e.A - e.B
		})
	defer patches.Reset()
	assert.Equal(t, -1, executeAdd(1, 2))
}
