package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// FileLogger 文件日志日志
type FileLogger struct {
	// Logger 底层
	*logrus.Logger

	// Logfile 日志文件
	Logfile string

	// Level 日志记录起始等级
	Level uint32
}

// NewFileLogger 生成新的日志对象
func NewFileLogger(filename string, level uint32) (*FileLogger, error) {
	l := &FileLogger{
		logrus.New(),
		filename,
		level,
	}

	err := l.Init()
	if nil != err {
		return nil, err
	}

	return l, nil
}

// Init 初始化
func (l *FileLogger) Init() error {
	fp, err := os.OpenFile(l.Logfile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if nil != err {
		return ErrOpenLogFileFail{l.Logfile, err}
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
func (l *FileLogger) Error(msg string, data map[string]interface{}) {
	l.WithFields(logrus.Fields(data)).Error(msg)
}

// Fatal shortcut
func (l *FileLogger) Fatal(msg string, data map[string]interface{}) {
	l.WithFields(logrus.Fields(data)).Fatal(msg)
}

// Panic shortcut
func (l *FileLogger) Panic(msg string, data map[string]interface{}) {
	l.WithFields(logrus.Fields(data)).Panic(msg)
}
