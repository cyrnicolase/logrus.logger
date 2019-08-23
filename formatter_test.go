package logger

import (
	"fmt"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestFormat(t *testing.T) {
	formatter := &MixFormatter{}
	tt := time.Now()
	level := "INFO"
	message := "Message"
	entry := logrus.WithFields(logrus.Fields{
		"foo": "bar",
	}).WithTime(tt)
	entry.Message = message
	entry.Level = logrus.Level(4)

	jsonData := `{"foo":"bar"}`
	testOutput := fmt.Sprintf("%s %s %s [%s] %s\n", tt.Format("2006-01-02 15:04:05"), level, "", message, jsonData)
	output, _ := formatter.Format(entry)

	if testOutput != string(output) {
		t.Errorf("期望：%v, 实际：%v\n", testOutput, string(output))
	}
}
