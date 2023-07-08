package types

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type HttpErr struct {
	Code    int
	Message string
	File    string // File where the error occurred
	Line    int    // Line number where the error occurred
}

func NewHttpErr(code int, message string) *HttpErr {
	_, file, line, _ := runtime.Caller(1)
	return &HttpErr{
		Code:    code,
		Message: message,
		File:    filepath.Base(file),
		Line:    line,
	}
}

func (httpErr *HttpErr) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s\nFile: %s, Line: %d", httpErr.Code, httpErr.Message, httpErr.File, httpErr.Line)
}
