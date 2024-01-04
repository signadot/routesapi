UID := $(shell id -u)
GID := $(shell id -g)

generate-openapi-doc:
	docker run --rm \
		-u ${UID}:${GID} \
		-v ${PWD}:/local \
		redocly/cli build-docs \
			-o /local/docs/http-api/index.html \
			/local/routes-openapi.yaml

