package cli

import (
	"regexp"
	"strings"
	"testing"

	"github.com/adriancarayol/ezdocker/internal/cli/mock"
	"github.com/adriancarayol/ezdocker/internal/tests"
)

func TestParseOptionsMinArgs(t *testing.T) {
	expected := `
Usage: ezd <option> <arguments>
Help: ezd help
`

	parser := New()
	out := tests.CaptureStdoutWrapper(parser.ParseOptions, nil)

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
	expected := "Invalid argument: a\n"

	mockClient := mock.DockerClient{}

	ConfigureCommands(mockClient)

	parser := New()
	out := tests.CaptureStdoutWrapper(parser.ParseOptions, []string{"test", "ls", "a"})

	if strings.Compare(out, expected) != 0 {
		t.Fatalf("Fail. Expected: %s, got: %s", expected, out)
	}
}
