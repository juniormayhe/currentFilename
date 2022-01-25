package currentFilename

import (
	"runtime"
	"strings"
)

func GetCurrentFileName() string {
	_, fullpath, _, _ := runtime.Caller(0)
	filename := strings.Split(fullpath, "/")

	return filename[len(filename)-1]
}
