package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/pkg/errors"

	"github.com/sliveryou/swag2md/pkg/builder"
	"github.com/sliveryou/swag2md/types"
)

const (
	methodGet            = "GET"
	responseSuccess      = 200
	schemaReferPrefix    = "#/definitions/"
	paramRequired        = "必填"
	paramNotRequired     = "非必填"
	reqFormat            = "| %s | %s | %s | %s | %s | %s |\n"
	respFormat           = "| %s | %s | %s | %s |\n"
	emsp                 = "&emsp;"
	casbinRuleAllow      = "allow" // 决策规则：允许
	casbinRuleDeny       = "deny"  // 决策规则：拒绝
	casbinPolicy         = "p, %s, %s, %s, %s\n"
	casbinWithDenyPolicy = "p, %s, %s, %s, %s, %s\n"
)

var replacer = strings.NewReplacer(
	" ", "-",
	".", "",
	"。", "",
	",", "-",
	"，", "-",
	"/", "",
	"(", "-",
	")", "",
	"[", "-",
	"]", "",
	"{", "-",
	"}", "",
	"（", "-",
	"）", "",
	"、", "-",
	";", "-",
	"；", "-",
)

// Parser Swagger解析器详情
type Parser struct {
	swagger   *types.Swagger
	keys      []string
	pathGroup map[string][]*types.PathInfo
	num       int
}

// NewParser 新建Swagger解析器
func NewParser(filePath string) (*Parser, error) {
	if filePath == "" {
		return nil, errors.New("illegal parser configure")
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.WithMessage(err, "os.ReadFile err")
	}

	swagger := &types.Swagger{}
	err = json.Unmarshal(content, swagger)
	if err != nil {
		return nil, errors.WithMessage(err, "json.Unmarshal err")
	}

	pathInfos := ParsePathInfos(swagger.Paths)
	keys, pathGroup := ParsePathGroup(pathInfos)

	return &Parser{
		swagger:   swagger,
		keys:      keys,
		pathGroup: pathGroup,
		num:       len(pathInfos),
	}, nil
}

// Build 构建文档
func (p *Parser) Build(title string) string {
	return p.BuildTitle(title) + p.BuildOverview() + p.BuildDetail()
}

// BuildTitle 构建标题
func (p *Parser) BuildTitle(title string) string {
	return fmt.Sprintf("# %s\n", title)
}

// BuildOverview 构建概览
func (p *Parser) BuildOverview() string {
	b := &strings.Builder{}
	b.WriteString(fmt.Sprintf("\n## 接口概览（总计%d个）\n", p.num))

	for _, key := range p.keys {
		b.WriteString(fmt.Sprintf("\n### %s\n", key))
		b.WriteString("\n| **路径** | **功能** | **请求方式** | **是否需要鉴权** |\n")
		b.WriteString("|---------|---------|-------------|-----------------|\n")

		if pathInfos, ok := p.pathGroup[key]; ok {
			for _, pi := range pathInfos {
				b.WriteString(fmt.Sprintf("| [%s](#%s) | %s | %s | %t |\n",
					pi.Path, replacer.Replace(pi.Summary), pi.Summary, pi.Method, pi.NeedAuth))
			}
		}
	}

	return b.String()
}

// BuildDetail 构建详情
func (p *Parser) BuildDetail() string {
	b := &strings.Builder{}
	b.WriteString("\n## 接口详情\n")

	for _, key := range p.keys {
		i := strings.Index(key, " ")
		if i == -1 {
			i = len(key)
		}

		b.WriteString(fmt.Sprintf("\n### %s\n", key[:i]))

		if pathInfos, ok := p.pathGroup[key]; ok {
			for _, pi := range pathInfos {
				b.WriteString(fmt.Sprintf("\n### %s\n", replacer.Replace(pi.Summary)))
				b.WriteString(fmt.Sprintf("\n[返回概览](#%s)\n\n---\n", replacer.Replace(key)))
				if pi.Method == methodGet {
					b.WriteString(fmt.Sprintf("\n%s %s\n", pi.Method, pi.Path))
				} else {
					b.WriteString(fmt.Sprintf("\n%s %s  \nContent-Type: %s\n", pi.Method, pi.Path, pi.Consume))
				}
				// 构建请求参数
				p.buildParameters(b, pi)
				// 构建响应参数
				p.buildResponses(b, pi)
			}
		}
	}

	return b.String()
}

// BuildCasbinPolicy 构建casbin规则
func (p *Parser) BuildCasbinPolicy(sub string, needDeny bool) string {
	b := &strings.Builder{}

	for _, key := range p.keys {
		if pathInfos, ok := p.pathGroup[key]; ok {
			for _, pi := range pathInfos {
				// rbac with deny policy
				// https://github.com/casbin/casbin/blob/master/examples/rbac_with_deny_policy.csv
				if needDeny {
					b.WriteString(fmt.Sprintf(casbinWithDenyPolicy, sub, pi.Path, pi.Method, casbinRuleAllow, pi.Summary))
				} else {
					b.WriteString(fmt.Sprintf(casbinPolicy, sub, pi.Path, pi.Method, pi.Summary))
				}
			}
		}
	}

	return b.String()
}

// buildParameters 构建请求参数
func (p *Parser) buildParameters(b *strings.Builder, pi *types.PathInfo) {
	if len(pi.Parameters) > 0 {
		eb := builder.NewExampleBuilder(pi.Path, false)

		b.WriteString("\n请求参数：\n")
		b.WriteString("\n| **来源** | **参数** | **描述** | **类型** | **约束** | **说明** |\n")
		b.WriteString("|----------|----------|----------|----------|----------|----------|\n")

		for _, param := range pi.Parameters {
			if param.Schema == nil {
				limit := paramNotRequired
				if param.Required {
					limit = paramRequired
				}

				b.WriteString(fmt.Sprintf(reqFormat, param.In, param.Name,
					convertDescription(param.Description, true),
					param.Type, limit,
					convertDescription(param.Description, false),
				))

				switch param.In {
				case types.ParameterInQuery:
					eb.AddQuery(param.Name, param.Type)
				case types.ParameterInFormData:
					eb.AddForm(param.Name, param.Type)
				case types.ParameterInHeader:
					eb.AddHeader(param.Name, param.Type)
				}
			} else {
				ref, isArray := getRefAndIsArray(param.Schema)
				eb.SetIsArray(isArray)
				dk := strings.TrimPrefix(ref, schemaReferPrefix)

				if d, ok := p.swagger.Definitions[dk]; ok && d.Type == types.SchemaTypeObject {
					ps := types.NewPropertySorter(d.Properties)
					rm := toMap(d.Required)
					for _, property := range ps {
						limit := paramNotRequired
						if _, ok := rm[property.Name]; ok {
							limit = paramRequired
						}

						b.WriteString(fmt.Sprintf(reqFormat, param.In, property.Name,
							convertDescription(property.Schema.Description, true),
							convertType(property.Schema), limit,
							convertDescription(property.Schema.Description, false),
						))

						pp := p.buildProperty(b, property, param.In, ref, 1)
						if pp == "" {
							eb.AddJSON(property.Name, property.Schema.Example, property.Schema.Type)
						} else {
							eb.AddJSONString(property.Name, pp, property.Schema.Type)
						}
					}
				}
			}
		}

		es := eb.String()
		if es != "" {
			b.WriteString("\n请求示例：\n")
			b.WriteString(es)
		}
	}
}

// buildResponses 构建响应参数
func (p *Parser) buildResponses(b *strings.Builder, pi *types.PathInfo) {
	if len(pi.Responses) > 0 {
		eb := builder.NewExampleBuilder(pi.Path, true)

		if resp, ok := pi.Responses[responseSuccess]; ok {
			var ref string
			if resp.Schema != nil && len(resp.Schema.AllOf) > 1 {
				for name, data := range resp.Schema.AllOf[1].Properties {
					ref, _ = getRefAndIsArray(data)
					eb.SetWrap(p.getWrap(resp.Schema.AllOf[0], name))
					break
				}
			} else if resp.Schema != nil && resp.Schema.Ref != "" {
				ref, _ = getRefAndIsArray(resp.Schema)
				eb.SetNeedWrap(false)
			}

			if ref != "" {
				b.WriteString(fmt.Sprintf("\n---\n\nContent-Type: %s\n", pi.Produce))
				b.WriteString("\n响应参数：\n")
				b.WriteString("\n| **参数** | **描述** | **类型** | **说明** |\n")
				b.WriteString("|----------|----------|----------|----------|\n")

				if resp.Schema != nil && resp.Schema.Type == types.SchemaTypeFile {
					b.WriteString(fmt.Sprintf(respFormat, resp.Schema.Type, "文件", resp.Schema.Type, resp.Description))
				}

				dk := strings.TrimPrefix(ref, schemaReferPrefix)
				if d, ok := p.swagger.Definitions[dk]; ok && d.Type == types.SchemaTypeObject {
					ps := types.NewPropertySorter(d.Properties)
					for _, property := range ps {
						b.WriteString(fmt.Sprintf(respFormat, property.Name,
							convertDescription(property.Schema.Description, true),
							convertType(property.Schema),
							convertDescription(property.Schema.Description, false),
						))

						pp := p.buildProperty(b, property, "", ref, 1)
						if pp == "" {
							eb.AddJSON(property.Name, property.Schema.Example, property.Schema.Type)
						} else {
							eb.AddJSONString(property.Name, pp, property.Schema.Type)
						}
					}
				}
			}
		}

		es := eb.String()
		if es != "" {
			b.WriteString("\n响应示例：\n")
			b.WriteString(es)
		}
	}
}

// buildProperty 构建属性参数
func (p *Parser) buildProperty(b *strings.Builder, pty *types.Property, paramIn, parentRef string, depth int) string {
	if pty.Schema.Type == types.SchemaTypeArray && pty.Schema.Items.Type != "" {
		return builder.GetArrayString(pty.Name, pty.Schema.Example, pty.Schema.Items.Type)
	}

	ref, isArray := getRefAndIsArray(pty.Schema)
	if parentRef == ref {
		return ""
	}

	if ref != "" {
		eb := builder.NewExampleBuilder("", false)
		dk := strings.TrimPrefix(ref, schemaReferPrefix)

		if d, ok := p.swagger.Definitions[dk]; ok && d.Type == types.SchemaTypeObject {
			ps := types.NewPropertySorter(d.Properties)
			rm := toMap(d.Required)

			for _, property := range ps {
				prefix := strings.Repeat(emsp, depth) + " "
				if paramIn != "" {
					limit := paramNotRequired
					if _, ok := rm[property.Name]; ok {
						limit = paramRequired
					}
					b.WriteString(fmt.Sprintf(reqFormat, paramIn, prefix+property.Name,
						convertDescription(property.Schema.Description, true),
						convertType(property.Schema), limit,
						convertDescription(property.Schema.Description, false),
					))
				} else {
					b.WriteString(fmt.Sprintf(respFormat, prefix+property.Name,
						convertDescription(property.Schema.Description, true),
						convertType(property.Schema),
						convertDescription(property.Schema.Description, false),
					))
				}

				pp := p.buildProperty(b, property, paramIn, ref, depth+1)
				if pp == "" {
					eb.AddJSON(property.Name, property.Schema.Example, property.Schema.Type)
				} else {
					eb.AddJSONString(property.Name, pp, property.Schema.Type)
				}
			}
		}

		jsonString := eb.GetJSON()
		if isArray {
			return fmt.Sprintf("[%s]", jsonString)
		}
		return jsonString
	}

	return ""
}

// getWrap 获取包装字符串
func (p *Parser) getWrap(s *types.Schema, dataName string) string {
	if s.Ref != "" {
		dk := strings.TrimPrefix(s.Ref, schemaReferPrefix)
		if d, ok := p.swagger.Definitions[dk]; ok && d.Type == types.SchemaTypeObject {
			b := &strings.Builder{}
			ps := types.NewPropertySorter(d.Properties, true)

			for _, property := range ps {
				if property.Name == dataName {
					b.WriteString(fmt.Sprintf("%q:%%s,", property.Name))
				} else {
					e := builder.GetInterfaceValue(property.Name, property.Schema.Type)
					if property.Schema.Example != nil {
						e = property.Schema.Example
					}
					example, _ := json.Marshal(e)

					b.WriteString(fmt.Sprintf("%q:%s,", property.Name, example))
				}
			}

			return fmt.Sprintf("{%s}", strings.TrimRight(b.String(), ","))
		}
	}

	return ""
}

// ParsePathInfos 解析路径信息列表
func ParsePathInfos(paths map[string]map[string]*types.Operation) []*types.PathInfo {
	var pathInfos []*types.PathInfo

	for path, methodOperation := range paths {
		for method, operation := range methodOperation {
			pi := &types.PathInfo{
				Method:     strings.ToUpper(method),
				Path:       path,
				Summary:    operation.Summary,
				Parameters: operation.Parameters,
				Responses:  operation.Responses,
			}

			if len(operation.Tags) > 0 {
				pi.Tag = strings.TrimSpace(operation.Tags[0])
			}

			if len(operation.Security) > 0 {
				pi.NeedAuth = true
			}

			if len(operation.Consumes) > 0 {
				pi.Consume = strings.TrimSpace(operation.Consumes[0])
			} else {
				pi.Consume = types.ApplicationJSON
			}

			if len(operation.Produces) > 0 {
				pi.Produce = strings.TrimSpace(operation.Produces[0])
			} else {
				pi.Produce = types.ApplicationJSON
			}

			pathInfos = append(pathInfos, pi)
		}
	}

	sort.Slice(pathInfos, func(i, j int) bool {
		pi, pj := pathInfos[i], pathInfos[j]
		if pi.Path != pj.Path {
			return pi.Path < pj.Path
		}
		li, lj := len(pi.Method), len(pj.Method)
		if li != lj {
			return li < lj
		}
		return pi.Method < pj.Method
	})

	return pathInfos
}

// ParsePathGroup 解析路径分组
func ParsePathGroup(pathInfos []*types.PathInfo) ([]string, map[string][]*types.PathInfo) {
	var keys []string
	m := make(map[string][]*types.PathInfo)

	for _, pi := range pathInfos {
		if _, ok := m[pi.Tag]; !ok {
			keys = append(keys, pi.Tag)
		}

		m[pi.Tag] = append(m[pi.Tag], pi)
	}
	sort.Strings(keys)

	return keys, m
}

// getRefAndIsArray 获取概要引用并判断其是否为数组
func getRefAndIsArray(schema *types.Schema) (ref string, isArray bool) {
	if schema.Type == types.SchemaTypeArray && schema.Items.Ref != "" {
		ref = schema.Items.Ref
		isArray = true
	} else if schema.Type == types.SchemaTypeObject || schema.Type == "" {
		ref = schema.Ref
		isArray = false
	}

	return
}

// convertType 转换类型
func convertType(schema *types.Schema) string {
	switch t := schema.Type; t {
	case "":
		return types.SchemaTypeObject
	case types.SchemaTypeArray:
		if it := schema.Items.Type; it != "" {
			return it + " " + types.SchemaTypeArray
		}
		return types.SchemaTypeObject + " " + types.SchemaTypeArray
	default:
		return t
	}
}

// convertDescription 转换描述
func convertDescription(d string, flag bool) string {
	left, right := d, ""
	start := strings.Index(d, "（")
	end := strings.LastIndex(d, "）")
	if start != -1 && end != -1 {
		left, right = d[:start], d[start+3:end]
	}

	if flag {
		return strings.TrimSpace(left)
	}
	return strings.TrimSpace(right)
}

// toMap slice转换map
func toMap(slice []string) map[string]struct{} {
	result := make(map[string]struct{}, len(slice))
	for _, v := range slice {
		result[v] = struct{}{}
	}

	return result
}
