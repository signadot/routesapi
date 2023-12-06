
generate-proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    	routes.proto
	perl -i -pe  s/SandboxId/SandboxID/g *.go
