package logger

import (
	"fmt"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestFormat(t *testing.T) {
	formatter := &MixFormatter{}
	t := time.Now()
	message := "Message"
	entry := logrus.WithFields(log.Fields{
		"foo": "bar",
	}).WithTime(t).Info(message)

	jsonData := `{"foo": "bar"}`
	testOutput := fmt.Sprintf("%s %s: [%s] %s\n", t, level, message, jsonData)
	output, err := formatter.Format(entry)

	if testOutput != output {
		t.Errorf("期望：%v, 实际：%v\n", testOutput, output)
	}
}
