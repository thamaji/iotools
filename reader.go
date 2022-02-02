package iotools

import (
	"bytes"
	"io"
)

type ReaderFunc func([]byte) (int, error)

func (r ReaderFunc) Read(p []byte) (int, error) {
	return r(p)
}

var Empty io.Reader = empty{}

type empty struct{}

func (empty) Read(p []byte) (int, error) {
	return 0, io.EOF
}

func NewSizeReader(r io.Reader) *SizeReader {
	return &SizeReader{r: r}
}

type SizeReader struct {
	r    io.Reader
	size int64
}

func (r *SizeReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	r.size += int64(n)
	return n, err
}

func (r *SizeReader) Size() int64 {
	return r.size
}

func NewPeekReader(r io.Reader) *PeekReader {
	return &PeekReader{r: r}
}

type PeekReader struct {
	r io.Reader
}

func (r *PeekReader) Read(p []byte) (int, error) {
	return r.r.Read(p)
}

func (r *PeekReader) Peek(size int) ([]byte, error) {
	buf := make([]byte, size)
	size, err := io.ReadFull(r.r, buf)
	if err != nil {
		return nil, err
	}
	buf = buf[:size]
	r.r = io.MultiReader(bytes.NewReader(buf), r.r)
	return buf, nil
}
