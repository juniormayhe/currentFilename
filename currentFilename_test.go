package currentFilename

import "testing"

func TestGetCurrentFilename(t *testing.T) {
	callerName := "currentFilename_test.go"

	filename := GetCurrentFileName()

	if filename != callerName {
		t.Errorf("Expected filename to be '%s', got %s", callerName, filename)
	}
}
