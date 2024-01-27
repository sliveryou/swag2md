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
		{k: "id1", v: nil, t: "string", expect: "[\"id\"]"},
		{k: "id2", v: []string{"a", "b", "c"}, t: "string", expect: "[\"a\",\"b\",\"c\"]"},
		{k: "id3", v: nil, t: "integer", expect: "[1]"},
		{k: "id4", v: []int{1, 2, 3}, t: "integer", expect: "[1,2,3]"},
		{k: "id5", v: nil, t: "object", expect: "null"},
	}

	for _, c := range cases {
		got := GetArrayString(c.k, c.v, c.t)
		assert.Equal(t, c.expect, got)
	}
}
