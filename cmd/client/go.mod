module vakond/client

go 1.13

require (
	github.com/jawher/mow.cli v1.1.0
	google.golang.org/grpc v1.27.1
	vakond/fileserver v0.0.0
)

replace vakond/fileserver => ../..
