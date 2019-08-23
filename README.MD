# 概述
借用[logrus](https://github.com/sirupsen/logrus)实现文件混合模式日志
增加部分快捷日志函数：
- func (l *FileLogger) Trace(msg string, data map[string]interface{})
- func (l *FileLogger) Debug(msg string, data map[string]interface{})
- func (l *FileLogger) Info(msg string, data map[string]interface{})
- func (l *FileLogger) Warn(msg string, data map[string]interface{})
- func (l *FileLogger) Warning(msg string, data map[string]interface{})
- func (l *FileLogger) Error(msg string, data map[string]interface{})
- func (l *FileLogger) Fatal(msg string, data map[string]interface{})
- func (l *FileLogger) Panic(msg string, data map[string]interface{})