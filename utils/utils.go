package utils

import (
	"oss/io"
	"strings"
	"time"
)

var directory string

func init() {
	conf, _ := io.ReadFile()
	directory = conf["directory"].(string)
}

func CreateFileName(name string) string {
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02")
	return directory + "/" + currentDate + "/" + name
}

func GetProtocol(endpoint string) string {
	protocolEnd := len("https://")
	return endpoint[:protocolEnd]
}

func GetDomain(endpoint string) string {
	protocolEnd := len("https://")
	return endpoint[protocolEnd:]
}

func ExtractFromURL(url string) string {
	index := strings.Index(url, directory)
	if index == -1 {
		return ""
	}
	return url[index:]
}

func Range(paths []string) []string {
	var result []string
	for _, str := range paths {
		result = append(result, ExtractFromURL(str))
	}
	return result
}
