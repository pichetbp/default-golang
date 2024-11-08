package helpers

import (
	"runtime"
	"strconv"
)

// methods
func SetFileLocation() string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	return StringBuild([]string{file, ":", strconv.Itoa(line)})

}
