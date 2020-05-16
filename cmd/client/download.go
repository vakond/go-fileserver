package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

func download(ver string) error {
	rpc, err := newRpc()
	if err != nil {
		return err
	}
	defer rpc.Close()

	stream, cancel, err := rpc.GetDownloadStream(ver)
	if err != nil {
		return rpc.Error(err)
	}
	defer cancel()

	packet, err := stream.Recv() // start downloading
	if err != nil {
		return err
	}

	filename := ver + ".zip" // arbitrary local filename
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer closeFile(file)

	var totalSize int
	for err == nil {
		n, e := file.Write(packet.Contents)
		if e != nil {
			return e
		}
		totalSize += n
		fmt.Printf("\r%d%% ", int64(math.Round(packet.Progress*100)))
		packet, err = stream.Recv()
	}
	if err != io.EOF {
		return err
	}

	fmt.Printf("Downloaded %d bytes into %s\n", totalSize, filename)

	return stream.CloseSend()
}

// closeFile dumps error from Close if any.
func closeFile(file io.Closer) {
	if err := file.Close(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
}
