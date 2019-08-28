package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	// FILE_MODE_SINGLE 单个文件模式
	FILE_MODE_SINGLE = iota

	// FILE_MODE_DAILY 按天分割
	FILE_MODE_DAILY
)

// FileMode 文件模式
type FileMode int

// FileLogger 文件日志日志
type FileLogger struct {
	// Logger 底层
	*logrus.Logger

	// Logfile 日志文件
	Logfile string

	// Level 日志记录起始等级
	Level uint32

	// 文件分割模式
	Mode FileMode
}

// NewFileLogger 生成新的日志对象
func NewFileLogger(filename string, level uint32, mode FileMode) (*FileLogger, error) {
	l := &FileLogger{
		logrus.New(),
		filename,
		level,
		mode,
	}

	err := l.Init()
	if nil != err {
		return nil, err
	}

	return l, nil
}

// Init 初始化
func (l *FileLogger) Init() error {
	fileName := l.Logfile
	if FILE_MODE_DAILY == l.Mode {
		fileName = fileName + time.Now().Format("2006-01-02")
	}
	fp, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if nil != err {
		return ErrOpenLogFileFail{fileName, err}
	}

	l.SetFormatter(&MixFormatter{})
	l.SetOutput(fp)
	l.SetLevel(logrus.Level(l.Level))

	return nil
}

// Trace shortcut
func (l *FileLogger) Trace(msg string, data map[string]interface{}) {
	l.WithFields(logrus.Fields(data)).Trace(msg)
}

// Debug shortcut
func (l *FileLogger) Debug(msg string, data map[string]interface{}) {
	l.WithFields(logrus.Fields(data)).Debug(msg)
}

// Info shortcut
func (l *FileLogger) Info(msg string, data map[string]interface{}) {
	l.WithFields(logrus.Fields(data)).Info(msg)
}

// Warn shortcut
func (l *FileLogger) Warn(msg string, data map[string]interface{}) {
	l.WithFields(logrus.Fields(data)).Warn(msg)
}

// Warning shortcut
func (l *FileLogger) Warning(msg string, data map[string]interface{}) {
	l.WithFields(logrus.Fields(data)).Warning(msg)
}

// Error shortcut
// receive a error interface
func (l *FileLogger) Error(msg interface{}, data map[string]interface{}) {
	entry := l.WithFields(logrus.Fields(data))
	if v, ok := msg.(string); ok {
		entry.Error(v)
	} else if v, ok := msg.(error); ok {
		entry.Error(v.Error())
	}
}

// Fatal shortcut
func (l *FileLogger) Fatal(msg string, data map[string]interface{}) {
	l.WithFields(logrus.Fields(data)).Fatal(msg)
}

// Panic shortcut
func (l *FileLogger) Panic(msg string, data map[string]interface{}) {
	l.WithFields(logrus.Fields(data)).Panic(msg)
}
