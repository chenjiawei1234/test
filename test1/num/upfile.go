package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	router.POST("/upfile", upFile)
	err := router.Run(":8089")
	if err != nil {
		log.Fatal(err.Error())
	}

}
func upFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// 创建文件
	out, err := os.Create("./test1/num" + file.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// 获取上传的文件
	f, err := file.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 将文件写入到我们新创建的文件中
	_, err = io.Copy(out, f)
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
