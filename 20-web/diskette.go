package main

import (
	"os"
)

type diskette struct {
	block *os.File
}

func (t *diskette) Write(p []byte) (n int, err error) {
	t.block.Truncate(0)
	t.block.Seek(0, 0)
	return t.block.Write(p)
}
