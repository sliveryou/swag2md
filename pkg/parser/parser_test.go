package parser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	filePath := "../../testdata/swagger.json"
	p, err := NewParser(filePath)
	assert.NoError(t, err)
	t.Logf("%+v", p)
	t.Logf(p.BuildTitle("接口文档"))
	t.Logf(p.BuildOverview())
	t.Logf(p.BuildDetail())

	outPath := "../../testdata/api.md"
	f, err := os.Create(outPath)
	assert.NoError(t, err)
	defer f.Close()

	_, _ = f.WriteString(p.BuildTitle("接口文档"))
	_, _ = f.WriteString(p.BuildOverview())
	_, _ = f.WriteString(p.BuildDetail())
}
