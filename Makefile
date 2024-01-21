.PHONY: fmt markdown policy

fmt:
	@find . -name '*.go' | xargs gofumpt -w -extra
	@find . -name '*.go' | xargs -n 1 -t goimports-reviser -rm-unused -set-alias -company-prefixes "github.com/sliveryou" -project-name "github.com/sliveryou/swag2md"

markdown:
	@swag2md -t "接口文档" -s swagger.json -o auto-gen-api.md

policy:
	@swag2md casbin -s swagger.json -o policy.csv --sub ADMIN --deny
