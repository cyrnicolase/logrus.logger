package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// 默认时间格式
var defaultTimestampForamt = "2006-01-02 15:04:05"

// MixFormatter 文本、Json混合模式
type MixFormatter struct {
	// 时间格式 默认使用 2006-02-03 15:04:05
	TimestampFormat string
}

// Format 实现Formatter 接口
func (f *MixFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestampFormat := f.TimestampFormat
	if "" == timestampFormat {
		timestampFormat = defaultTimestampForamt
	}

	if entry.Time.IsZero() {
		entry.Time = time.Now()
	}
	t := entry.Time.Format(timestampFormat)
	level := strings.ToUpper(entry.Level.String())
	message := entry.Message
	funcVal := ""
	if entry.HasCaller() {
		funcVal = fmt.Sprintf("%s#%d:%s", entry.Caller.File, entry.Caller.Line, entry.Caller.Function)
	}

	b := &bytes.Buffer{}
	encoder := json.NewEncoder(b)
	if err := encoder.Encode(entry.Data); nil != err {
		return nil, fmt.Errorf("结构化数据转JSON失败, 原因：%v", err)
	}

	jsonData := strings.Trim(string(b.Bytes()), "\n")
	result := fmt.Sprintf("%s %s %s [%s] %s\n", t, level, funcVal, message, jsonData)

	return []byte(result), nil
}
