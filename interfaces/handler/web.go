package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type WebAppHandle interface {
	IndexHandler(c *gin.Context)
	PlayPageHander(c *gin.Context)
	UploadHandler(c *gin.Context)
}

type webAppHandle struct {

}

func(w *webAppHandle) IndexHandler(c *gin.Context){
	c.HTML(http.StatusOK,"index.html","")
}
func(w *webAppHandle) PlayPageHander(c *gin.Context){
	c.HTML(http.StatusOK,"PlayVideo.html","")
}
func(w *webAppHandle) UploadHandler(c *gin.Context){
	c.HTML(http.StatusOK,"upload.html",nil)
}

func NewWebAppHandler()WebAppHandle{
	return &webAppHandle{}
}