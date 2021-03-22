package usecase

import (
	"os"

	"video-distribution-site-golang-layerd-arch/domain"
	"video-distribution-site-golang-layerd-arch/domain/repository"
)

type VideoUseCase interface {
	GetAll(page int64) (*[]domain.VideoInfo,error)
	GetVideoURI(videoID string)
	Delete(videoID string)(error)
	Upload(title string,fileIns *os.File,filesize int64)(*domain.VideoSrc,error)
}

type videoUseCase struct {
	videoRepositoty repository.VideoRepository
}

func NewVideoUseCase(repositoty repository.VideoRepository) VideoUseCase{
	return &videoUseCase{videoRepositoty: repositoty}
}

func (v *videoUseCase) GetAll(page int64) (*[]domain.VideoInfo, error) {
	return v.videoRepositoty.GetAll(page)
}

func (v *videoUseCase) GetVideoURI(videoID string) {
	panic("implement me")
}

func (v *videoUseCase) Delete(videoID string) error {
	return v.videoRepositoty.Delete(videoID)
}

func (v *videoUseCase) Upload(title string,fileIns *os.File,filesize int64) (*domain.VideoSrc,error) {
	videoFileStruct,err:=domain.NewVideo(title,filesize,fileIns)
	if err != nil {
		return nil,err
	}
	v.videoRepositoty.Save(videoFileStruct)
	return videoFileStruct, nil
}

var _ VideoUseCase = &videoUseCase{}

