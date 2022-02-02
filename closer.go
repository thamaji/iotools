package iotools

type CloserFunc func() error

func (c CloserFunc) Close() error {
	return c()
}

var NopCloser = nopCloser{}

type nopCloser struct{}

func (nopCloser) Close() error {
	return nil
}
