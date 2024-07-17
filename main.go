package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"oss/io"
	"oss/sdk"
)

type Path struct {
	Path string `json:"path"`
}

type Paths struct {
	Paths []string `json:"paths"`
}

var (
	port     string
	endpoint string
	bucket   string
)

func init() {
	fmt.Println("init")
	conf, err := io.ReadFile()
	if err != nil {
		fmt.Println("Failed to read config file:", err)
	}
	endpoint = conf["Endpoint"].(string)
	bucket = conf["Bucket"].(string)
	port = conf["port"].(string)
}

func upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"path": handler(sdk.Upload(file)),
	})
}

func handler(uri string) string {
	prefix := fmt.Sprintf("%s%s.%s", getProtocol(endpoint), bucket, getDomain(endpoint))
	return prefix + "/" + uri
}

func getProtocol(endpoint string) string {
	protocolEnd := len("https://")
	return endpoint[:protocolEnd]
}

func getDomain(endpoint string) string {
	protocolEnd := len("https://")
	return endpoint[protocolEnd:]
}

func single(c *gin.Context) {
	b, _ := c.GetRawData()
	body := Path{}
	_ = json.Unmarshal(b, &body)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s deleted!", sdk.Single(body.Path)),
	})
}

func multiple(c *gin.Context) {
	b, _ := c.GetRawData()
	body := Paths{}
	_ = json.Unmarshal(b, &body)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s deleted!", sdk.Multiple(body.Paths)),
	})
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/upload", upload)
	r.DELETE("/single", single)
	r.DELETE("/multiple", multiple)
	err := r.Run(":" + port)
	if err != nil {
		return
	}
}
