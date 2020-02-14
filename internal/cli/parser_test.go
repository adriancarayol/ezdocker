package cli

import (
	"reflect"
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

func TestParseParameters(t *testing.T) {
	expected := []string{"ls", "a"}

	mockClient := mock.DockerClient{}

	ConfigureCommands(mockClient)

	parser := New()

	out, _ := parser.parseParameters([]string{"test", "ls", "a"})

	if reflect.DeepEqual(expected, out) {
		t.Fatalf("Fail. Expected: %s, got: %s", expected, out)
	}
}
