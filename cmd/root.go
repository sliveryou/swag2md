package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/gookit/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/sliveryou/swag2md/pkg/markdown"
	"github.com/sliveryou/swag2md/pkg/parser"
)

const (
	// codeFailure 失败错误码
	codeFailure = 1
	// buildVersion swag2md版本
	buildVersion = "1.0.1"
)

var (
	// title 生成的markdown文本标题
	title string
	// swaggerFile swagger文件路径
	swaggerFile string
	// outputFile 输出文件路径
	outputFile string
)

var rootCmd = &cobra.Command{
	Use:   "swag2md",
	Short: "Convert Swagger 2.0 JSON document to markdown document",
	Long:  "Convert JSON documents that comply with Swagger 2.0 into more user-friendly markdown format interface documents",
	Example: "  swag2md -t \"接口文档\" -s swagger.json -o api.md\n" +
		"  swag2md casbin -s swagger.json -o policy.csv --sub ADMIN --deny",
	RunE: rootFunc,
}

func init() {
	rootCmd.Version = fmt.Sprintf("%s %s/%s", buildVersion, runtime.GOOS, runtime.GOARCH)

	rootCmd.Flags().StringVarP(&title, "title", "t", "接口文档", "the title of output markdown text")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "auto-gen-api.md", "the file to store output markdown text")
	rootCmd.PersistentFlags().StringVarP(&swaggerFile, "swagger", "s", "swagger.json", "the swagger.json file")
}

func rootFunc(_ *cobra.Command, args []string) error {
	p, err := parser.NewParser(swaggerFile)
	if err != nil {
		return errors.WithMessage(err, "parser.NewParser err")
	}

	f, err := os.Create(outputFile)
	if err != nil {
		return errors.WithMessage(err, "os.Create err")
	}
	defer f.Close()

	out, err := markdown.Process("", []byte(p.Build(title)), nil)
	if err != nil {
		return errors.WithMessage(err, "markdown.Process err")
	}

	_, err = f.Write(out)
	if err != nil {
		return errors.WithMessage(err, "outputFile.Write err")
	}

	return nil
}

// Execute executes the root command and its subcommands.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(color.Red.Render(err.Error()))
		os.Exit(codeFailure)
	}
}
