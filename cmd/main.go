package main

import (
	"github.com/gin-gonic/gin"
	"video-distribution-site-golang-layerd-arch/infrastructure"
	"video-distribution-site-golang-layerd-arch/interfaces/handler"
	"video-distribution-site-golang-layerd-arch/usecase"
)

func main() {
	Videorepository:=infrastructure.NewVideoRepository("")
	videoUsecase:=usecase.NewVideoUseCase(Videorepository)
	videoHanler:=handler.NewVideoHandler(videoUsecase)
	
	webHanler:=handler.NewWebAppHandler()
	router:=gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static","./static")
	// 簡単のためにvideos0
	router.Static("/videos", "./videos")


	router.GET("/", webHanler.IndexHandler)
	router.GET("/play", webHanler.PlayPageHander)
	router.GET("/upload",webHanler.UploadHandler )

	router.POST("/api/index",videoHanler.HandleVideoGetAll)
	router.POST("/api/upload",videoHanler.HandleUpload)

}
