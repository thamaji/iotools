package iotools

import "io"

type WriterFunc func([]byte) (int, error)

func (w WriterFunc) Write(p []byte) (int, error) {
	return w(p)
}

func NewSizeWriter(w io.Writer) *SizeWriter {
	return &SizeWriter{w: w}
}

type SizeWriter struct {
	w    io.Writer
	size int64
}

func (w *SizeWriter) Write(p []byte) (int, error) {
	n, err := w.w.Write(p)
	w.size += int64(n)
	return n, err
}

func (w *SizeWriter) Size() int64 {
	return w.size
}
