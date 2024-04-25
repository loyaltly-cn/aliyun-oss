package sdk

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	"os"
	"oss/io"
	"oss/utils"
)

var bucket *oss.Bucket

func init() {

	conf, _ := io.ReadFile()
	client, err := oss.New(conf["Endpoint"].(string), conf["AccessKeyId"].(string), conf["AccessKeySecret"].(string))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	b, err := client.Bucket(conf["Bucket"].(string))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	bucket = b
}

func Upload(fileHeader *multipart.FileHeader) string {

	file, err := fileHeader.Open()
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	uri := utils.CreateFileName(fileHeader.Filename)

	err = bucket.PutObject(uri, file)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	return uri
}

func Single(name string) bool {
	err := bucket.DeleteObject(name)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	return true
}

func Multiple(paths []string) bool {
	_, err := bucket.DeleteObjects(paths,
		oss.DeleteObjectsQuiet(true))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	return true
}
