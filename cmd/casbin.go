package cmd

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/sliveryou/swag2md/pkg/parser"
)

var (
	// outputCsv 输出csv文件路径
	outputCsv string
	// subjectName 访问实体名称
	subjectName string
	// needDeny 是否需要拒绝选项
	needDeny bool
)

var casbinCmd = &cobra.Command{
	Use:     "casbin",
	Short:   "Generate casbin policy file",
	Example: "  swag2md casbin -s swagger.json -o policy.csv --sub ADMIN --deny",
	RunE:    casbinFunc,
}

func init() {
	casbinCmd.Flags().StringVarP(&outputCsv, "output", "o", "policy.csv", "the file to store output casbin policy")
	casbinCmd.Flags().StringVarP(&subjectName, "sub", "", "ADMIN", "the casbin policy subject name")
	casbinCmd.Flags().BoolVarP(&needDeny, "deny", "", false, "whether need the deny field")

	rootCmd.AddCommand(casbinCmd)
}

func casbinFunc(_ *cobra.Command, _ []string) error {
	p, err := parser.NewParser(swaggerFile)
	if err != nil {
		return errors.WithMessage(err, "parser.NewParser err")
	}

	if err := os.WriteFile(outputCsv, []byte(p.BuildCasbinPolicy(subjectName, needDeny)), writeFilePerm); err != nil {
		return errors.WithMessage(err, "os.WriteFile err")
	}

	return nil
}
