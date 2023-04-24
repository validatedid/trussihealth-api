package ipfs

import (
	"bytes"
	shell "github.com/ipfs/go-ipfs-api"
	"io"
)

type IPFSClientWrapper struct {
	wrapper *shell.Shell
}

func NewIPFSClientWrapper(client *shell.Shell) *IPFSClientWrapper {
	return &IPFSClientWrapper{wrapper: client}
}

func (i *IPFSClientWrapper) Add(data *bytes.Reader) (string, error) {
	// add file to IPFS
	hash, err := i.wrapper.Add(data)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (i *IPFSClientWrapper) Cat(path string) (io.ReadCloser, error) {
	data, err := i.wrapper.Cat(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}
