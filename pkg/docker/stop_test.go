package docker

import (
	"github.com/adriancarayol/ezdocker/pkg/cli/mock"
	"github.com/adriancarayol/ezdocker/pkg/tests"
	"regexp"
	"strings"
	"testing"
)

func TestStopContainersCommand_HandleHelp(t *testing.T) {
	expected := `
stop: <option>s
-a Stop all running containers
`
	mockClient := mock.NotEmptyDockerClient{}
	client := New(mockClient)

	cmd := StopContainersCommand{Docker: client}
	out := tests.CaptureStdoutWrapper(func() {
		cmd.Handle("-b", "-b")
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

func TestStopContainersCommand_HandleStopAllContainers(t *testing.T) {
	expected := "StoppingcontainerwithID:testingID..."

	mockClient := mock.NotEmptyDockerClient{}
	client := New(mockClient)

	cmd := StopContainersCommand{Docker: client}
	out := tests.CaptureStdoutWrapper(func() {
		cmd.Handle("-a")
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

func TestStopContainersCommand_HandleStopContainers(t *testing.T) {
	expected := "StoppingcontainerwithID:testingID..."

	mockClient := mock.NotEmptyDockerClient{}
	client := New(mockClient)

	cmd := StopContainersCommand{Docker: client}
	out := tests.CaptureStdoutWrapper(func() {
		cmd.Handle("testingID")
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

func TestStopContainersCommand_HandleErrorStopContainers(t *testing.T) {
	expected := "StoppingcontainerwithID:testingID...ErrorstoppingcontainerwithID:testingID..."

	mockClient := mock.ErrorDockerClient{}
	client := New(mockClient)

	cmd := StopContainersCommand{Docker: client}
	out := tests.CaptureStdoutWrapper(func() {
		cmd.Handle("testingID")
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

func TestStopContainersCommand_HandleErrorStopNotContainers(t *testing.T) {
	expected := "Not running containers."

	mockClient := mock.DockerClient{}
	client := New(mockClient)

	cmd := StopContainersCommand{Docker: client}
	out := tests.CaptureStdoutWrapper(func() {
		cmd.Handle("-a")
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