package main

import "github.com/cyrnicolase/logger"

func main() {
	log := logger.NewFileLogger("./test.log", 5)
	log.Init()

	log.Info("写一个Info日志", map[string]interface{}{
		"foo": "bar",
		"数字":  12345676,
	})
}
