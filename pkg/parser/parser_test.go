package parser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sliveryou/swag2md/pkg/markdown"
)

const (
	writeFilePerm = 0o666
)

func TestNewParser(t *testing.T) {
	_, err := NewParser("")
	require.EqualError(t, err, "illegal parser configure")

	_, err = NewParser("not/exist/swagger/file")
	require.EqualError(t, err, "os.ReadFile err: open not/exist/swagger/file: no such file or directory")

	_, err = NewParser("parser.go")
	require.EqualError(t, err, "json.Unmarshal err: invalid character 'p' looking for beginning of value")

	p, err := NewParser("../../examples/swagger.json")
	require.NoError(t, err)
	assert.NotNil(t, p)
}

func TestParser_Build(t *testing.T) {
	title := "示例接口文档"
	filePath := "../../examples/swagger.json"
	p, err := NewParser(filePath)
	require.NoError(t, err)

	doc := p.Build(title)
	t.Log(doc)

	outPath := "../../examples/swagger.md"
	out, err := markdown.Process("", []byte(doc), nil)
	require.NoError(t, err)

	err = os.WriteFile(outPath, out, writeFilePerm)
	require.NoError(t, err)
}

func TestParser_BuildCasbinPolicy(t *testing.T) {
	sub := "ADMIN"
	filePath := "../../examples/swagger.json"
	p, err := NewParser(filePath)
	require.NoError(t, err)

	policy := p.BuildCasbinPolicy(sub, false)
	t.Log(policy)

	outPath := "../../examples/policy.csv"
	err = os.WriteFile(outPath, []byte(policy), writeFilePerm)
	require.NoError(t, err)
}

func TestParser_BuildCasbinPolicyDeny(t *testing.T) {
	sub := "ADMIN"
	filePath := "../../examples/swagger.json"
	p, err := NewParser(filePath)
	require.NoError(t, err)

	policy := p.BuildCasbinPolicy(sub, true)
	t.Log(policy)

	outPath := "../../examples/policy-deny.csv"
	err = os.WriteFile(outPath, []byte(policy), writeFilePerm)
	require.NoError(t, err)
}
