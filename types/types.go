package types

import (
	"sort"
)

// PathInfo 路径详情
type PathInfo struct {
	Tag        string
	Method     string
	Path       string
	Summary    string
	NeedAuth   bool
	Consume    string
	Produce    string
	Parameters []*Parameter
	Responses  map[int]*Response
}

// Property 属性字段详情
type Property struct {
	Name   string
	Schema *Schema
}

// PropertySorter 属性字段排序器
type PropertySorter []*Property

// Swap 实现 Swap 方法
func (ps PropertySorter) Swap(i, j int) { ps[i], ps[j] = ps[j], ps[i] }

// Len 实现 Len 方法
func (ps PropertySorter) Len() int { return len(ps) }

// Less 实现 Less 方法
func (ps PropertySorter) Less(i, j int) bool {
	if ps[i].Schema.XOrder != ps[j].Schema.XOrder {
		return ps[i].Schema.XOrder < ps[j].Schema.XOrder
	}

	return ps[i].Name < ps[j].Name
}

// NewPropertySorter 新建属性字段排序器
func NewPropertySorter(m map[string]*Schema) PropertySorter {
	var ps PropertySorter

	for name, schema := range m {
		ps = append(ps, &Property{Name: name, Schema: schema})
	}
	sort.Sort(ps)

	return ps
}
