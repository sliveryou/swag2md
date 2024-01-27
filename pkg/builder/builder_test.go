package builder

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewExampleBuilder(t *testing.T) {
	eb := NewExampleBuilder("/example/api", true)
	assert.NotNil(t, eb)
	assert.Equal(t, "{\"trace_id\":\"a1b2c3d4e5f6g7h8\",\"code\":0,\"msg\":\"ok\",\"data\":%s}", eb.wrap)
}

func TestExampleBuilder_String(t *testing.T) {
	eb := NewExampleBuilder("/example/api", true)

	eb.AddQuery("id", "integer")
	eb.AddQuery("name", "string")
	eb.AddForm("school", "string")
	eb.AddForm("score", "number")
	eb.AddHeader("Auth", "string")
	eb.AddHeader("Key", "string")

	out := eb.String()
	fmt.Println(out)
	assert.Equal(t, "\n```\nHeader:\nAuth: Auth\nKey: Key\nQuery:\n/example/api?id=1&name=name\nForm Data:\nschool: school\nscore: 1.0\n```\n", out)
}

func TestExampleBuilder_String2(t *testing.T) {
	eb := NewExampleBuilder("/example/api", true)

	eb.AddQuery("id", "integer")
	eb.AddQuery("name", "string")
	eb.AddHeader("Auth", "string")
	eb.AddHeader("Key", "string")
	eb.AddJSON("name", nil, "string")
	eb.AddJSON("school", "test-school", "string")
	eb.AddJSON("score", 90, "integer")
	eb.AddJSONString("number", "", "integer")
	eb.AddJSONString("phone", "123456", "integer")

	out := eb.String()
	fmt.Println(out)
	assert.Equal(t, "\n```\nHeader:\nAuth: Auth\nKey: Key\nQuery:\n/example/api?id=1&name=name\n```\n\n```json\n{\n  \"trace_id\": \"a1b2c3d4e5f6g7h8\",\n  \"code\": 0,\n  \"msg\": \"ok\",\n  \"data\": {\n    \"name\": \"name\",\n    \"school\": \"test-school\",\n    \"score\": 90,\n    \"number\": 1,\n    \"phone\": 123456\n  }\n}\n```\n", out)
}

func TestExampleBuilder_String3(t *testing.T) {
	eb := NewExampleBuilder("/example/api", true)

	eb.AddQuery("id", "integer")
	eb.AddQuery("name", "string")
	eb.AddHeader("Auth", "string")
	eb.AddHeader("Key", "string")
	eb.AddJSON("name", nil, "string")
	eb.AddJSON("school", "test-school", "string")
	eb.AddJSON("score", 90, "integer")
	eb.AddJSONString("number", "", "integer")
	eb.AddJSONString("phone", "123456", "integer")
	eb.SetNeedWrap(false)

	out := eb.String()
	fmt.Println(out)
	assert.Equal(t, "\n```\nHeader:\nAuth: Auth\nKey: Key\nQuery:\n/example/api?id=1&name=name\n```\n\n```json\n{\n  \"name\": \"name\",\n  \"school\": \"test-school\",\n  \"score\": 90,\n  \"number\": 1,\n  \"phone\": 123456\n}\n```\n", out)
}

func TestExampleBuilder_String4(t *testing.T) {
	eb := NewExampleBuilder("/example/api", true)

	eb.AddQuery("id", "integer")
	eb.AddQuery("name", "string")
	eb.AddHeader("Auth", "string")
	eb.AddHeader("Key", "string")
	eb.AddJSON("name", nil, "string")
	eb.AddJSON("school", "test-school", "string")
	eb.AddJSON("score", 90, "integer")
	eb.AddJSONString("number", "", "integer")
	eb.AddJSONString("phone", "123456", "integer")
	eb.SetIsArray(true)

	out := eb.String()
	fmt.Println(out)
	assert.Equal(t, "\n```\nHeader:\nAuth: Auth\nKey: Key\nQuery:\n/example/api?id=1&name=name\n```\n\n```json\n{\n  \"trace_id\": \"a1b2c3d4e5f6g7h8\",\n  \"code\": 0,\n  \"msg\": \"ok\",\n  \"data\": [\n    {\n      \"name\": \"name\",\n      \"school\": \"test-school\",\n      \"score\": 90,\n      \"number\": 1,\n      \"phone\": 123456\n    }\n  ]\n}\n```\n", out)
}

func TestExampleBuilder_GetJSON(t *testing.T) {
	eb := NewExampleBuilder("/example/api", true)

	eb.AddQuery("id", "integer")
	eb.AddJSON("name", nil, "string")
	eb.AddJSON("school", "test-school", "string")
	eb.AddJSON("score", 90, "integer")
	eb.AddJSONString("number", "", "integer")
	eb.AddJSONString("phone", "123456", "integer")

	out := eb.GetJSON()
	fmt.Println(out)
	assert.Equal(t, "{\"name\":\"name\",\"school\":\"test-school\",\"score\":90,\"number\":1,\"phone\":123456}", out)
}

func TestGetArrayString(t *testing.T) {
	defaultKey := "id"
	cases := []struct {
		k      string
		v      interface{}
		t      string
		expect string
	}{
		{k: defaultKey, v: nil, t: "string", expect: "[\"id\"]"},
		{k: defaultKey, v: []string{"a", "b", "c"}, t: "string", expect: "[\"a\",\"b\",\"c\"]"},
		{k: defaultKey, v: nil, t: "integer", expect: "[1]"},
		{k: defaultKey, v: []int{1, 2, 3}, t: "integer", expect: "[1,2,3]"},
		{k: defaultKey, v: nil, t: "object", expect: "null"},
		{k: defaultKey, v: nil, t: "number", expect: "[1.0]"},
		{k: defaultKey, v: nil, t: "boolean", expect: "[false]"},
		{k: defaultKey, v: nil, t: "file", expect: "null"},
		{k: defaultKey, v: nil, t: "unknown", expect: "null"},
		{k: defaultKey, v: []float64{1.1, 2.2, 3.3}, t: "number", expect: "[1.1,2.2,3.3]"},
		{k: "id2", v: nil, t: "string", expect: "[\"id2\"]"},
		{k: "id3", v: nil, t: "string", expect: "[\"id3\"]"},
	}

	for _, c := range cases {
		got := GetArrayString(c.k, c.v, c.t)
		assert.Equal(t, c.expect, got)
	}
}

func TestGetInterfaceValue(t *testing.T) {
	defaultKey := "id"
	cases := []struct {
		k      string
		t      string
		expect interface{}
	}{
		{k: defaultKey, t: "string", expect: "id"},
		{k: defaultKey, t: "integer", expect: 1},
		{k: defaultKey, t: "number", expect: 1.0},
		{k: defaultKey, t: "boolean", expect: false},
		{k: defaultKey, t: "object", expect: nil},
		{k: defaultKey, t: "file", expect: "(file)"},
		{k: defaultKey, t: "array", expect: nil},
		{k: defaultKey, t: "unknown", expect: nil},
		{k: "id2", t: "string", expect: "id2"},
		{k: "id3", t: "string", expect: "id3"},
	}

	for _, c := range cases {
		got := GetInterfaceValue(c.k, c.t)
		assert.Equal(t, c.expect, got)
	}
}

func Test_getStringValue(t *testing.T) {
	defaultKey := "id"
	cases := []struct {
		k      string
		t      string
		expect string
	}{
		{k: defaultKey, t: "string", expect: defaultKey},
		{k: defaultKey, t: "integer", expect: "1"},
		{k: defaultKey, t: "number", expect: "1.0"},
		{k: defaultKey, t: "boolean", expect: "false"},
		{k: defaultKey, t: "object", expect: "id"},
		{k: defaultKey, t: "file", expect: "(file)"},
		{k: defaultKey, t: "array", expect: defaultKey},
		{k: defaultKey, t: "unknown", expect: defaultKey},
		{k: "id2", t: "string", expect: "id2"},
		{k: "id3", t: "string", expect: "id3"},
	}

	for _, c := range cases {
		got := getStringValue(c.k, c.t)
		assert.Equal(t, c.expect, got)
	}
}

func Test_getJSONStringValue(t *testing.T) {
	defaultKey := "id"
	cases := []struct {
		k      string
		t      string
		expect string
	}{
		{k: defaultKey, t: "string", expect: `"id"`},
		{k: defaultKey, t: "integer", expect: "1"},
		{k: defaultKey, t: "number", expect: "1.0"},
		{k: defaultKey, t: "boolean", expect: "false"},
		{k: defaultKey, t: "object", expect: "null"},
		{k: defaultKey, t: "file", expect: `"(file)"`},
		{k: defaultKey, t: "array", expect: "null"},
		{k: defaultKey, t: "unknown", expect: "null"},
		{k: "id2", t: "string", expect: `"id2"`},
		{k: "id3", t: "string", expect: `"id3"`},
	}

	for _, c := range cases {
		got := getJSONStringValue(c.k, c.t)
		assert.Equal(t, c.expect, got)
	}
}
