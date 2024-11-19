//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest --config=types.cfg.yaml openapi.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest --config=server.cfg.yaml openapi.core.v1

package coreopenapi
