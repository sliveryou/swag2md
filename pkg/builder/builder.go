package builder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/sliveryou/swag2md/types"
)

// ExampleBuilder 示例构建器
type ExampleBuilder struct {
	path     string
	needWrap bool
	isArray  bool
	query    url.Values
	form     string
	header   string
	json     string
	wrap     string
}

// NewExampleBuilder 新建示例构建器
func NewExampleBuilder(path string, needWrap bool) *ExampleBuilder {
	return &ExampleBuilder{
		path:     path,
		needWrap: needWrap,
		query:    make(url.Values),
		// 默认包装字符串
		wrap: "{\"trace_id\":\"a1b2c3d4e5f6g7h8\",\"code\":0,\"msg\":\"ok\",\"data\":%s}",
	}
}

// SetNeedWrap 设置是否需要封装响应
func (eb *ExampleBuilder) SetNeedWrap(needWrap bool) {
	eb.needWrap = needWrap
}

// SetWrap 设置封装响应
func (eb *ExampleBuilder) SetWrap(wrap string) {
	if wrap != "" && strings.Contains(wrap, "%s") {
		eb.wrap = wrap
	}
}

// SetIsArray 设置参数是否为数组
func (eb *ExampleBuilder) SetIsArray(isArray bool) {
	eb.isArray = isArray
}

// String 构建示例字符串
func (eb *ExampleBuilder) String() string {
	b := &strings.Builder{}

	if eb.header != "" || len(eb.query) > 0 || eb.form != "" {
		b.WriteString("\n```\n")

		if eb.header != "" {
			b.WriteString("Header:\n" + eb.header)
		}

		if len(eb.query) > 0 {
			b.WriteString(fmt.Sprintf("Query:\n%s?%s\n", eb.path, eb.query.Encode()))
		}

		if eb.form != "" {
			b.WriteString("Form Data:\n" + eb.form)
		}

		b.WriteString("```\n")
	}

	if eb.json != "" {
		b.WriteString("\n```json\n")

		s := eb.GetJSON()

		if eb.isArray {
			s = fmt.Sprintf("[%s]", s)
		}

		if eb.needWrap {
			s = fmt.Sprintf(eb.wrap, s)
		}

		out := &bytes.Buffer{}
		_ = json.Indent(out, []byte(s), "", "  ")

		b.WriteString(out.String() + "\n```\n")
	}

	return b.String()
}

// AddQuery 添加Query类型参数
func (eb *ExampleBuilder) AddQuery(k, t string) {
	eb.query.Add(k, getStringValue(k, t))
}

// AddForm 添加Form类型参数
func (eb *ExampleBuilder) AddForm(k, t string) {
	eb.form += fmt.Sprintf("%s: %s\n", k, getStringValue(k, t))
}

// AddHeader 添加Header类型参数
func (eb *ExampleBuilder) AddHeader(k, t string) {
	eb.header += fmt.Sprintf("%s: %s\n", k, getStringValue(k, t))
}

// AddJSON 添加JSON类型参数
func (eb *ExampleBuilder) AddJSON(k string, v interface{}, t string) {
	var e interface{}
	if v != nil {
		e = v
	} else {
		e = GetInterfaceValue(k, t)
	}
	example, _ := json.Marshal(e)
	eb.json += fmt.Sprintf("%q:%s,", k, string(example))
}

// AddJSONString 添加JSON字符串类型参数
func (eb *ExampleBuilder) AddJSONString(k, v, t string) {
	var e string
	if v != "" {
		e = v
	} else {
		e = getJSONStringValue(k, t)
	}
	eb.json += fmt.Sprintf("%q:%s,", k, e)
}

// GetJSON 获取JSON类型参数
func (eb *ExampleBuilder) GetJSON() string {
	return fmt.Sprintf("{%s}", strings.TrimRight(eb.json, ","))
}

// GetArrayString 获取array string类型的实例值
func GetArrayString(k string, v interface{}, t string) string {
	if v != nil {
		e, _ := json.Marshal(v)
		return string(e)
	}

	switch t {
	case types.SchemaTypeString:
		return fmt.Sprintf("[%q]", k)
	case types.SchemaTypeInteger:
		return "[1]"
	case types.SchemaTypeNumber:
		return "[1.0]"
	case types.SchemaTypeBoolean:
		return "[false]"
	default:
		return "null"
	}
}

// GetInterfaceValue 获取interface{}类型的示例值
func GetInterfaceValue(k, t string) interface{} {
	var v interface{}

	switch t {
	case types.SchemaTypeString:
		v = k
	case types.SchemaTypeInteger:
		v = 1
	case types.SchemaTypeNumber:
		v = 1.0
	case types.SchemaTypeBoolean:
		v = false
	case types.SchemaTypeFile:
		v = "(file)"
	case types.SchemaTypeObject:
		v = nil
	case types.SchemaTypeArray:
		v = nil
	}

	return v
}

// getStringValue 获取string类型的示例值
func getStringValue(k, t string) string {
	var v string

	switch t {
	case types.SchemaTypeInteger:
		v = "1"
	case types.SchemaTypeNumber:
		v = "1.0"
	case types.SchemaTypeBoolean:
		v = "false"
	case types.SchemaTypeFile:
		v = "(file)"
	default:
		v = k
	}

	return v
}

// getJSONStringValue 获取json string类型的示例值
func getJSONStringValue(k, t string) string {
	var v string

	switch t {
	case types.SchemaTypeString:
		v = fmt.Sprintf("%q", k)
	case types.SchemaTypeInteger:
		v = "1"
	case types.SchemaTypeNumber:
		v = "1.0"
	case types.SchemaTypeBoolean:
		v = "false"
	case types.SchemaTypeFile:
		v = "\"(file)\""
	default:
		v = "null"
	}

	return v
}
