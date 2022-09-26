.DEFAULT_GOAL := build

build: tools/openapi-generator-cli clean generate fmt

clean:
	rm -rf api/

generate:
	openapi-generator-cli generate -i "https://data.commercelayer.app/schemas/openapi-no-ref.json" \
		--generator-name go \
		--output api \
		--ignore-file-override .openapi-generator-ignore \
		--config config.yml

fmt:
	go fmt ./...

tools/openapi-generator-cli:
	npm install -g @openapitools/openapi-generator-cli