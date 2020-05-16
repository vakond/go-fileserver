all: fileserver client

fileserver:
	go mod tidy
	go fmt ./...
	go vet ./...
	rm -f api/*.pb.go
	go generate
	go install -trimpath

client:
	cd cmd/client; \
	go mod tidy; \
	go fmt ./...; \
	go vet ./...; \
	go install -trimpath
