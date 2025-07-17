REDOCLY_VERSION := latest
OAPI_CODEGEN_VERSION := latest

.PHONY: openapi

openapi:
	@echo "Generating code from OpenAPI specifications..."
	@echo "Creating dist directory if it does not exist..."
	mkdir -p dist
	@echo "Bundling OpenAPI specification with Redocly CLI..."
	docker run --user=$(shell id -u):$(shell id -g) --rm -v$(PWD):/app -w /app redocly/cli:$(REDOCLY_VERSION) bundle openapi/openapi.yml -o dist/openapi.yml
	@echo "Installing oapi-codegen..."
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@$(OAPI_CODEGEN_VERSION)
	@echo "Generating Go code from OpenAPI specification..."
	oapi-codegen -generate types,client,server,spec,skip-prune -package generated -o openapi/generated.go dist/openapi.yml
