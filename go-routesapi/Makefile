UID := $(shell id -u)
GID := $(shell id -g)

generate-proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		-I .. \
    	../routes.proto
	perl -i -pe  s/SandboxId/SandboxID/g *.go

generate-openapi-models:
	cp ../routes-openapi.yaml .
	docker run --rm \
		-u ${UID}:${GID} \
		-v ${PWD}:/local \
		-e GO_POST_PROCESS_FILE="/usr/local/bin/gofmt -w" \
		openapitools/openapi-generator-cli generate \
			-i /local/routes-openapi.yaml \
			-g go \
			-o /local/models \
			--global-property models,modelDocs=false,supportingFiles=utils.go \
			--additional-properties=withGoMod=false,packageName=models
	rm routes-openapi.yaml
	rm -rf ./models/api
	rm -rf ./models/.openapi-generator
