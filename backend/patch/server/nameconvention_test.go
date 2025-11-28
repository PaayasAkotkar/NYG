package server

import (
	"regexp"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestNamingConvention(t *testing.T) {
	name := "user_6bdf845e-f656-4a80-a3fe-dc678a33d254"
	re := regexp.MustCompile("_|-|[^a-zA-Z]")
	newName := re.ReplaceAllString(name, "")[:8]
	assert.Equal(t, newName, "userbdfe")
}
