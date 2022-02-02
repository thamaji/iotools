package iotools

import "io"

func WriteCloser(w io.Writer, c io.Closer) io.WriteCloser {
	return struct {
		io.Writer
		io.Closer
	}{
		w, c,
	}
}

func ReadCloser(r io.Reader, c io.Closer) io.ReadCloser {
	return struct {
		io.Reader
		io.Closer
	}{
		r, c,
	}
}

func ReadWriteCloser(r io.Reader, w io.Writer, c io.Closer) io.WriteCloser {
	return struct {
		io.Reader
		io.Writer
		io.Closer
	}{
		r, w, c,
	}
}
