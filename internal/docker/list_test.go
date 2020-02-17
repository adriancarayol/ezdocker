package docker

import (
	"github.com/adriancarayol/ezdocker/internal/cli/mock"
	"github.com/adriancarayol/ezdocker/internal/tests"
	"regexp"
	"strings"
	"testing"
)

func TestPrintContainersCommand_HandleEmpty(t *testing.T) {
	expected := "•ID:testingID•IMAGE:•STATUS:-•COMMAND:•PORTS:•IP:192.0.0.0•Publicport:90•Privateport:80•Protocol:TCP"

	mockClient := mock.NotEmptyDockerClient{}
	client := New(mockClient)

	cmd := PrintContainersCommand{Docker: client}
	out := tests.CaptureStdoutWrapper(func() {
		cmd.Handle()
	}, nil)

	re := regexp.MustCompile(`\r?\n`)
	expectedSanitized := re.ReplaceAllString(expected, " ")
	outSanitized := re.ReplaceAllString(out, " ")

	expectedSanitized = strings.Replace(expectedSanitized, " ", "", -1)
	outSanitized = strings.Replace(outSanitized, " ", "", -1)

	if strings.Compare(expectedSanitized, outSanitized) != 0 {
		t.Fatalf("Expected: %v, got: %v", expectedSanitized, outSanitized)
	}
}

func TestPrintContainersCommand_HandleMinimal(t *testing.T) {
	expected := "•ID:testingID•IMAGE:•STATUS:-"

	mockClient := mock.NotEmptyDockerClient{}
	client := New(mockClient)

	cmd := PrintContainersCommand{Docker: client}
	out := tests.CaptureStdoutWrapper(func() {
		cmd.Handle("-m")
	}, nil)

	re := regexp.MustCompile(`\r?\n`)
	expectedSanitized := re.ReplaceAllString(expected, " ")
	outSanitized := re.ReplaceAllString(out, " ")

	expectedSanitized = strings.Replace(expectedSanitized, " ", "", -1)
	outSanitized = strings.Replace(outSanitized, " ", "", -1)

	if strings.Compare(expectedSanitized, outSanitized) != 0 {
		t.Fatalf("Expected: %v, got: %v", expectedSanitized, outSanitized)
	}
}

func TestPrintContainersCommand_HandleHelp(t *testing.T) {
	expected := `
ls: <option>s
-m Minimal information (id, image, status)
`
	mockClient := mock.NotEmptyDockerClient{}
	client := New(mockClient)

	cmd := PrintContainersCommand{Docker: client}
	out := tests.CaptureStdoutWrapper(func() {
		cmd.Handle("b", "b")
	}, nil)

	re := regexp.MustCompile(`\r?\n`)
	expectedSanitized := re.ReplaceAllString(expected, " ")
	outSanitized := re.ReplaceAllString(out, " ")

	expectedSanitized = strings.Replace(expectedSanitized, " ", "", -1)
	outSanitized = strings.Replace(outSanitized, " ", "", -1)

	if strings.Compare(expectedSanitized, outSanitized) != 0 {
		t.Fatalf("Expected: %v, got: %v", expectedSanitized, outSanitized)
	}
}

func TestPrintContainersCommand_HandleError(t *testing.T) {
	expected := "There's no containers running."

	mockClient := mock.ErrorDockerClient{}
	client := New(mockClient)

	cmd := PrintContainersCommand{Docker: client}
	out := tests.CaptureStdoutWrapper(func() {
		cmd.Handle("-m")
	}, nil)

	re := regexp.MustCompile(`\r?\n`)
	expectedSanitized := re.ReplaceAllString(expected, " ")
	outSanitized := re.ReplaceAllString(out, " ")

	expectedSanitized = strings.Replace(expectedSanitized, " ", "", -1)
	outSanitized = strings.Replace(outSanitized, " ", "", -1)

	if strings.Compare(expectedSanitized, outSanitized) != 0 {
		t.Fatalf("Expected: %v, got: %v", expectedSanitized, outSanitized)
	}
}
