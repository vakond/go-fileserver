package main

import (
	"fmt"
)

func versions() error {
	rpc, err := newRpc()
	if err != nil {
		return err
	}
	defer rpc.Close()

	resp, err := rpc.Versions()
	if err != nil {
		return rpc.Error(err)
	}

	for _, ver := range resp.Version {
		fmt.Println(ver)
	}

	return nil
}
