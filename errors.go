package logger

import "fmt"

// ErrOpenLogFileFail 打开日志文件失败
type ErrOpenLogFileFail struct {
	// 日志文件名
	Filename string

	// 上级错误
	Err error
}

// IsErrOpenLogFileFail 是否为打开日志文件失败错误
func IsErrOpenLogFileFail(err error) bool {
	_, ok := err.(ErrOpenLogFileFail)
	return ok
}

func (e ErrOpenLogFileFail) Error() string {
	return fmt.Sprintf("打开日志文件失败, filename: %s, err: %v", e.Filename, e.Err)
}
