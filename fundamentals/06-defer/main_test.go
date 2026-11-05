package main

import (
	"errors"
	"testing"
)

type failingWriter struct {
	writeErr, closeErr error
	closed             bool
}

func (w *failingWriter) Write(p []byte) (int, error) { return 0, w.writeErr }
func (w *failingWriter) Close() error                { w.closed = true; return w.closeErr }
func TestCleanupPreservesWriteError(t *testing.T) {
	writeErr, closeErr := errors.New("write"), errors.New("close")
	w := &failingWriter{writeErr: writeErr, closeErr: closeErr}
	if err := writeContent(w); !errors.Is(err, writeErr) {
		t.Fatal(err)
	}
	if !w.closed {
		t.Fatal("recurso não fechado após erro")
	}
}
func TestCloseErrorReturned(t *testing.T) {
	cause := errors.New("close")
	good := &closeFails{cause: cause}
	if err := writeContent(good); !errors.Is(err, cause) {
		t.Fatal(err)
	}
}

type closeFails struct{ cause error }

func (w *closeFails) Write(p []byte) (int, error) { return len(p), nil }
func (w *closeFails) Close() error                { return w.cause }
