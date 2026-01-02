package service

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

const (
	ParamError   = "param error"
	SaveFileFail = "save file fail"
)

type FileOp struct {
	DefaultDir string
}

var fileOp = FileOp{"./data"}

func NewFileOp(fileDir string) *FileOp {
	return &FileOp{
		DefaultDir: fileDir,
	}
}

func (f *FileOp) LoadService(g *gin.RouterGroup) error {
	g.POST("/upload", f.upload)
	g.POST("/list", f.list)
	g.POST("/download", f.download)
	return nil
}

func (f *FileOp) upload(c *gin.Context) {
	container := c.PostForm("container")
	key := c.PostForm("key")
	dstPath := path.Join(f.DefaultDir, container, key)
	dir := path.Dir(dstPath)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		c.String(http.StatusInternalServerError, "create dir error")
		return
	}

	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	log.Printf("dstPath:%s", dstPath)
	log.Printf("dir:%s", dir)
	err = c.SaveUploadedFile(file, dstPath)
	if err != nil {
		log.Printf("[save file fail, errMsg:%s]", err.Error())
		c.String(http.StatusInternalServerError, fmt.Sprintf("error_msg:%s", SaveFileFail))
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("%s upload", file.Filename))
}

func (f *FileOp) list(c *gin.Context) {

}

type downloadReq struct {
	Container string `json:"container" binding:"required"`
	Key       string `json:"key" binding:"required"`
}

func (f *FileOp) download(c *gin.Context) {
	var req downloadReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error_msg:%s", ParamError))
		return
	}
	dst := path.Join(f.DefaultDir, req.Container, req.Key)
	log.Printf("dst:%s", dst)
	c.File(dst)
}
