package main

import (
	"encoding/json"
	"fmt"
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

func upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%s' uploaded!", sdk.Upload(file)),
	})
}

func single(c *gin.Context) {
	b, _ := c.GetRawData()
	body := Path{}
	_ = json.Unmarshal(b, &body)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%s' deleted!", sdk.Single(body.Path)),
	})
}

func multiple(c *gin.Context) {
	b, _ := c.GetRawData()
	body := Paths{}
	_ = json.Unmarshal(b, &body)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%s' deleted!", sdk.Multiple(body.Paths)),
	})
}

func main() {
	conf, _ := io.ReadFile()
	port := ":" + conf["port"].(string)
	r := gin.Default()

	r.POST("/upload", upload)
	r.DELETE("/single", single)
	r.DELETE("/multiple", multiple)
	err := r.Run(port)
	if err != nil {
		return
	}

}
