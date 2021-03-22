package repository

import (
	"video-distribution-site-golang-layerd-arch/domain"
)

type VideoRepository interface {
	Save(videoSrc *domain.VideoSrc) (error)
	GetAll(offset int64) (*[]domain.VideoInfo,error)
	Delete(videoID string)error
	FindByVideoID(videoID string)*domain.VideoInfo
}
