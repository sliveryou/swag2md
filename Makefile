.PHONY: porxy tidy fmt lint test pre-commit markdown policy

proxy:
	@go env -w GO111MODULE="on"
	@go env -w GOPROXY="https://goproxy.cn,direct"

tidy:
	@go mod tidy -e -v

fmt:
	@find . -name '*.go' | xargs gofumpt -w -extra
	@find . -name '*.go' | xargs -n 1 -t goimports-reviser -rm-unused -set-alias -company-prefixes "github.com/sliveryou" -project-name "github.com/sliveryou/swag2md"
	@find . -name '*.sh' | xargs shfmt -w -s -i 2 -ci -bn -sr

lint:
	@golangci-lint run ./...

test:
	@go test ./...

pre-commit: tidy fmt lint test

markdown:
	@swag2md -t "接口文档" -s swagger.json -o auto-gen-api.md

policy:
	@swag2md casbin -s swagger.json -o policy.csv --sub ADMIN --deny
