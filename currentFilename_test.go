package currentFilename

import "testing"

func TestGetCurrentFilename(t *testing.T) {
	filename := GetCurrentFileName()

	if filename != "currentFilename.go" {
		t.Errorf("Expected filename to be 'currentFilename.go', got %s", filename)
	}
}
