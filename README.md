# swag2md

[![Github License](https://img.shields.io/github/license/sliveryou/swag2md.svg?style=flat)](https://github.com/sliveryou/swag2md/blob/master/LICENSE)
[![Go Doc](https://godoc.org/github.com/sliveryou/swag2md?status.svg)](https://pkg.go.dev/github.com/sliveryou/swag2md)
[![Go Report](https://goreportcard.com/badge/github.com/sliveryou/swag2md)](https://goreportcard.com/report/github.com/sliveryou/swag2md)
[![Github Latest Release](https://img.shields.io/github/release/sliveryou/swag2md.svg?style=flat)](https://github.com/sliveryou/swag2md/releases/latest)
[![Github Latest Tag](https://img.shields.io/github/tag/sliveryou/swag2md.svg?style=flat)](https://github.com/sliveryou/swag2md/tags)
[![Github Stars](https://img.shields.io/github/stars/sliveryou/swag2md.svg?style=flat)](https://github.com/sliveryou/swag2md/stargazers)


swag2md 是一个可以将符合 Swagger 2.0 的 JSON 文档转化成较为友好的 Markdown 格式的接口文档的工具。

## 安装

使用如下命令下载并安装 swag2md 工具：

```bash
# 如果 go 版本在 1.16 以前，使用如下命令安装：
$ GO111MODULE=on go get -u github.com/sliveryou/swag2md@latest

# 如果 go 版本在 1.16 及以后，使用如下命令安装：
$ GO111MODULE=on go install github.com/sliveryou/swag2md@latest
```

## swag2md 命令行接口

```bash
$ swag2md -h
swag2md 命令的用法:
  -o string
        解析输出的 markdown 文件的名称 (默认 "auto-gen-api.md")
  -s string
        待解析的 swagger.json 文件 (默认 "swagger.json")
  -t string
        解析输出的 markdown 文件内容的标题 (默认 "接口文档")

例子:
  swag2md -t "接口文档" -s swagger.json -o api.md
```

**PS：**

- swagger.json 文件可以参考使用 `swag init` 命令生成
- swag 工具文档：[https://github.com/swaggo/swag/blob/master/README_zh-CN.md](https://github.com/swaggo/swag/blob/master/README_zh-CN.md)
