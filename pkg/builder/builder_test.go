package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetArrayString(t *testing.T) {
	cases := []struct {
		k      string
		v      interface{}
		t      string
		expect string
	}{
		{k: "id", v: nil, t: "string", expect: "[\"id\"]"},
		{k: "id", v: []string{"a", "b", "c"}, t: "string", expect: "[\"a\",\"b\",\"c\"]"},
		{k: "id", v: nil, t: "integer", expect: "[1]"},
		{k: "id", v: []int{1, 2, 3}, t: "integer", expect: "[1,2,3]"},
		{k: "id", v: nil, t: "object", expect: "null"},
	}

	for _, c := range cases {
		got := GetArrayString(c.k, c.v, c.t)
		assert.Equal(t, c.expect, got)
	}
}
