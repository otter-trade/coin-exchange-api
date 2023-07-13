package util

import (
	"runtime"
	"strings"
)

func GetFuncName() string {
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		name := runtime.FuncForPC(pc).Name()
		idx := strings.LastIndex(name, "/")
		if idx >= 0 {
			name = name[idx+1:]
		}
		return name
	}
	return ""
}
