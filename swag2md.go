package main

import (
	"flag"
	"log"
	"os"

	"github.com/sliveryou/swag2md/pkg/markdown"
	"github.com/sliveryou/swag2md/pkg/parser"
)

var (
	// title 生成的markdown文本标题
	title = flag.String("t", "接口文档", "the title of output markdown text")
	// swaggerFile swagger文件路径
	swaggerFile = flag.String("s", "swagger.json", "the swagger.json file")
	// outputFile 输出文件路径
	outputFile = flag.String("o", "auto-gen-api.md", "the file to store output markdown text")
)

func main() {
	flag.Parse()

	p, err := parser.NewParser(*swaggerFile)
	if err != nil {
		log.Println(err)
		return
	}

	f, err := os.Create(*outputFile)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	out, err := markdown.Process("", []byte(p.Build(*title)), nil)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = f.Write(out)
	if err != nil {
		log.Println(err)
		return
	}
}
