package file

import (
	"os"
	"testing"
)

// TestLogger_Log for testing log file
func TestLogger_Log(t *testing.T) {
	loggerTest := Create()
	loggerTest.Config()
	loggerTest.Log("Hi")

	if _, err := os.Stat(loggerTest.Destination); err != nil {
		t.Errorf("%s: %s", "Expected to log file exists but got", err)
	}

	_ = os.Remove(loggerTest.Destination)
}
