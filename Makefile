all: fileserver client

fileserver:
	rm -f api/*.pb.go
	go generate
	go mod tidy
	go fmt ./...
	go vet ./...
	go install -trimpath

client:
	cd cmd/client; \
	go mod tidy; \
	go fmt ./...; \
	go vet ./...; \
	go install -trimpath
