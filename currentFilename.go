package currentFilename

import (
	"runtime"
	"strings"
)

func GetCurrentFileName() string {
	// caller 0 is this package file name
	// caller 1 is the external caller that uses this package
	_, caller, _, _ := runtime.Caller(1)
	filename := strings.Split(caller, "/")

	return filename[len(filename)-1]
}
