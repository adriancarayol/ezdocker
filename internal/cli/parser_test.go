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

	oldArgs := os.Args
	os.Args = []string{}

	parser.ParseOptions()
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

	re := regexp.MustCompile(`\r?\n`)
	expectedSanitized := re.ReplaceAllString(expected, " ")
	outSanitized := re.ReplaceAllString(out, " ")

	expectedSanitized = strings.Replace(expectedSanitized, " ", "", -1)
	outSanitized = strings.Replace(outSanitized, " ", "", -1)

	if strings.Compare(expectedSanitized, outSanitized) != 0 {
		t.Fatalf("Fail: Expected: %s, got: %s\n", expectedSanitized, outSanitized)
	}
}

func TestParseOptionsInvalidArg(t *testing.T) {
	expected := "Invalid option: a\n"

	// Init()

	parser := New()
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	oldArgs := os.Args
	os.Args = []string{"test", "ls", "a"}

	parser.ParseOptions()
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

	if strings.Compare(out, expected) != 0 {
		t.Fatalf("Fail. Expected: %s, got: %s", expected, out)
	}
}
