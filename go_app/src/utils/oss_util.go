package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// UploadFromURL 从一个URL上传文件到S3
func UploadFromURL(url string) string {
	res, err := http.Get(url)
	if err != nil {
		panic("download image error" + err.Error())
	}
	file, err := ioutil.ReadAll(res.Body)
	fileReader := bytes.NewReader(file)

	key := createOssKey()
	if strings.Contains(url, "ooo.0o0.ooo") {
		keyList := strings.Split(url, "/")
		key = "blog/" + keyList[3] + keyList[4] + "/" + keyList[6]
	}

	if os.Getenv("APP_ENV") != "prod" {
		key = os.Getenv("APP_ENV") + "/" + key
	}
	key = key + ".jpg"
	DoUpload(key, fileReader)
	return key
}

func DoUpload(key string, fileReader io.Reader) {
	client, err := oss.New(os.Getenv("OSS_HOST"), os.Getenv("OSS_ACCESS_KEY"), os.Getenv("OSS_SECRET_KEY"))
	if err != nil {
		panic("oss client build error" + err.Error())
	}

	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	if err != nil {
		panic("oss client bucket error" + err.Error())
	}
	err = bucket.PutObject(key, fileReader)
	if err != nil {
		panic("put object error" + err.Error())
	} else {
		fmt.Print(os.Getenv("END_POINT") + key)
	}
}

func createOssKey() string {
	return "blog/" + time.Now().Format("200601") + "/" + strconv.FormatInt(time.Now().Unix(), 10)
}
