package main

import (
	"errors"
	"fmt"

	"github.com/cyrnicolase/logger"
)

func main() {
	log, _ := logger.NewFileLogger("./test.log", 5, 0)

	log.Debug("写一个Info日志", map[string]interface{}{
		"foo": "bar",
		"数字":  12345676,
	})

	log1, err := logger.NewFileLogger("/root/test1.log", 4, 1)

	if nil != err {
		fmt.Println(err)
		return
	}

	log1.Debug("写一个Debug 日志", map[string]interface{}{
		"debug": "DEBUG",
	})

	// 测试Error接收error情况
	log.Error(errors.New("写入一个错误接口对象"), nil)
}
