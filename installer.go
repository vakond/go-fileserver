package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"sort"

	"vakond/fileserver/api"
)

// fileserver is a gRPC handler.
type fileserver struct{}

// Versions handles call 'versions'.
func (i *fileserver) Versions(_ context.Context, _ *api.Request) (*api.VersionsResponse, error) {
	var resp api.VersionsResponse
	for ver := range config.versions {
		resp.Version = append(resp.Version, ver)
	}
	sort.Strings(resp.Version) // TODO: sort as semvers, not strings
	return &resp, nil
}

// Download handles call 'download'.
// Download breaks the file contents into chunks and sends them one by one.
func (i *fileserver) Download(req *api.Request, stream api.Fileserver_DownloadServer) error {
	const chunkSize = 1024 // bytes

	filename, found := config.versions[req.Version]
	if !found {
		return fmt.Errorf("unsupported version '%s'", req.Version)
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer closeFile(file)

	// Get file size to calculate progress
	info, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := info.Size()
	var currentSize int64
	var size int

	// Read file and send it's contents down to the client
	reader := bufio.NewReader(file)
	chunk := make([]byte, chunkSize)
	var packet api.DownloadResponse
	for err == nil {
		size, err = reader.Read(chunk)
		if size > 0 {
			packet.Contents = chunk[:size] // size may differ from chunkSize
			currentSize += int64(size)
			packet.Progress = float64(currentSize) / float64(fileSize)
			if e := stream.Send(&packet); e != nil {
				return e
			}
			packet.Index++
		}
	}
	if err != io.EOF {
		return err
	}

	return nil
}
