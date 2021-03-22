package domain

import (
	"errors"
	"github.com/google/uuid"
	"os"
)

const MAX_FILE_SIZE = 1024 * 1024*1024*1024 * 1024*1024

type VideoFileName struct {
	fileName string
}

type URL struct {
	URL string
}

type VideoID struct {
	ID  string
}

type VideoCount struct{
	Count int
}

type VideoInfo struct {
	Title   string `json:"Name"`
	VideoID string `json:"videoid"`
	Count   int64  `json:"count"`
	Link    string `json:"link"`
}

type VideoSrc struct{
	VideoInfo
	VideoIns *os.File
}


func NewVideo(title string,filesize int64,fileIns *os.File)(*VideoSrc,error){
	if filesize>MAX_FILE_SIZE{
		return nil,errors.New("UPLOADED FILE IS TOO LARGE")
	}
	if title==""{
		return nil,errors.New("Title is empty")
	}
	uid,err:=uuid.NewUUID()
	if err != nil {
		return nil,err
	}
	v:=&VideoSrc{
		VideoInfo: VideoInfo{
			Title:   title,
			VideoID: uid.String(),
			Count:   0,
			Link:    "play/"+uid.String(),
		},
		VideoIns:  fileIns,
	}
	return v,nil
}
