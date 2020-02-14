package tests

import (
	"bytes"
	"io"
	"os"
)

func CaptureStdoutWrapper(f func(), fakeArgs []string) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	oldArgs := os.Args

	if fakeArgs != nil {
		os.Args = fakeArgs
	} else {
		os.Args = []string{}
	}

	f()

	os.Args = oldArgs

	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	out := <-outC

	return out
}
