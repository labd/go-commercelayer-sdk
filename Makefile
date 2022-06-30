.DEFAULT_GOAL := build

build: tools/openapi-generator-cli generate fmt

generate:
	openapi-generator-cli generate -i "https://data.commercelayer.app/schemas/openapi.json" \
		--generator-name go \
		--output api \
		--config config.yml

fmt:
	go fmt ./...

tools/openapi-generator-cli:
	npm install -g @openapitools/openapi-generator-cli