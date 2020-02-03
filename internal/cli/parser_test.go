package cli

import (
	"bytes"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestParseOptionsMinArgs(t *testing.T) {
	expected := `
Usage: ezd <option> <arguments>
Help: ezd help
`

	parser := New()
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	os.Args = []string{}

	parser.ParseOptions()

	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	out := <-outC

	re := regexp.MustCompile(`\r?\n`)
	expectedSanitized := re.ReplaceAllString(expected, " ")
	outSanitized := re.ReplaceAllString(out, " ")

	expectedSanitized = strings.Replace(expectedSanitized, " ", "", -1)
	outSanitized = strings.Replace(outSanitized, " ", "", -1)

	if strings.Compare(expectedSanitized, outSanitized) != 0 {
		t.Fatalf("Fail: Expected: %s, got: %s\n", expectedSanitized, outSanitized)
	}
}
