package handler

import (
	"net/http"
	"os"
	"strconv"
	"video-distribution-site-golang-layerd-arch/usecase"
)
import "github.com/gin-gonic/gin"

type VideoHandler interface {
	HandleVideoGetAll(c *gin.Context)
	HandleGetVideoURI(c *gin.Context)
	HandleUpload(c *gin.Context)
}

type videoHandler struct {
	videoUsecase usecase.VideoUseCase
}

func (v *videoHandler) HandleVideoGetAll(c *gin.Context) {
	page,err:=strconv.ParseInt(c.Query("page"),10,64)
	if err != nil {
		return
	}
	allvideo,err:=v.videoUsecase.GetAll(page)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"videoInfos":allvideo,
	})
}

func (v *videoHandler) HandleGetVideoURI(c *gin.Context) {
	panic("implement me")
}

func (v *videoHandler) HandleUpload(c *gin.Context) {
	// file size取得とos.fileの取得
	fileSize := c.Request.Header.Get("content-length")
	size, err := strconv.ParseInt(fileSize,10,64)
	if  err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
		})
	}
	file, err := c.FormFile("videoFile")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	tmp,err:=file.Open()
	f:=tmp.(*os.File)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail upload",
		})
	}

	// titleの取得
	title:=c.PostForm("title")
	result,err:=v.videoUsecase.Upload(title,f,size)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail upload",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"result":result,
	})
}

func NewVideoHandler(usecase usecase.VideoUseCase)VideoHandler{
	return &videoHandler{usecase}
}