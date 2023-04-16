package miniredis

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
)

func TestMiniRedis(t *testing.T) {
	s := miniredis.RunT(t)
	err := s.Set("foo", "bar")
	assert.NoError(t, err)
	s.CheckGet(t, "foo", "bar")
	s.SetTTL("foo", 10*time.Second)
	s.FastForward(11 * time.Second)
	if s.Exists("foo") {
		t.Fatal("'foo' should not have existed anymore")
	}
}
