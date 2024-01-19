package types

// ParameterIn 参数来源
const (
	ParameterInQuery    = "query"
	ParameterInFormData = "formData"
	ParameterInHeader   = "header"
)

// SchemaType 概要类型
const (
	SchemaTypeString  = "string"
	SchemaTypeInteger = "integer"
	SchemaTypeNumber  = "number"
	SchemaTypeBoolean = "boolean"
	SchemaTypeFile    = "file"
	SchemaTypeObject  = "object"
	SchemaTypeArray   = "array"
)

const (
	// ApplicationJSON 应用类型：json
	ApplicationJSON = "application/json"
)

// Swagger swagger详情
type Swagger struct {
	BasePath    string                           `json:"basePath"`
	Info        *Info                            `json:"info"`
	Paths       map[string]map[string]*Operation `json:"paths"`
	Definitions map[string]*Definition           `json:"definitions"`
}

// Info 项目详情
type Info struct {
	Description string `json:"description"`
	Title       string `json:"title"`
	Version     string `json:"version"`
}

// Operation 操作详情
type Operation struct {
	Consumes   []string              `json:"consumes"`
	Produces   []string              `json:"produces"`
	Tags       []string              `json:"tags"`
	Summary    string                `json:"summary"`
	Security   []map[string][]string `json:"security"`
	Parameters []*Parameter          `json:"parameters"`
	Responses  map[int]*Response     `json:"responses"`
}

// Parameter 参数详情
type Parameter struct {
	Description string  `json:"description"`
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	In          string  `json:"in"`
	Required    bool    `json:"required"`
	Schema      *Schema `json:"schema"`
}

// Response 响应详情
type Response struct {
	Description string  `json:"description"`
	Schema      *Schema `json:"schema"`
}

// Schema 概要详情
type Schema struct {
	Ref         string             `json:"$ref"`
	Description string             `json:"description"`
	Type        string             `json:"type"`
	XOrder      string             `json:"x-order"`
	Properties  map[string]*Schema `json:"properties"`
	AllOf       []*Schema          `json:"allOf"`
	Items       *Schema            `json:"items"`
	Example     interface{}        `json:"example"`
}

// Definition 定义详情
type Definition struct {
	Type       string             `json:"type"`
	Required   []string           `json:"required"`
	Properties map[string]*Schema `json:"properties"`
}
