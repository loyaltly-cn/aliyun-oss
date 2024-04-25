package utils

import (
	"oss/io"
	"time"
)

func CreateFileName(name string) string {
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02")
	conf, _ := io.ReadFile()
	return conf["directory"].(string) + "/" + currentDate + "/" + name
}
