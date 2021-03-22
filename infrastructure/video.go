package infrastructure

import (
	"database/sql"
	"io"
	"os"
	
	"video-distribution-site-golang-layerd-arch/domain/repository"
	"video-distribution-site-golang-layerd-arch/domain"
)

type videoRepository struct{
	saveDir string
	db *sql.DB
}

const tableName= "video"
// video
/*
	videoID Title count
*/
func (v *videoRepository) Save(file *domain.VideoSrc) (error) {

	v.db.Exec("INSERT INTO ? values()")
	src := file.VideoIns
	defer src.Close()

	out, err := os.Create(v.saveDir+file.VideoID)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func (v *videoRepository) CountUp(videoID string){

}

func (v *videoRepository) GetAll(offset int64) (*[]domain.VideoInfo, error) {
	rows,err:=v.db.Query("select * from ? ORDER BY `count` ASC LIMIT 100 OFFSET ?",tableName,offset)
	if err != nil {
		return nil, err
	}

	ansList:=make([]domain.VideoInfo,0,100)
	for rows.Next() {
		v:=domain.VideoInfo{}
		if err := rows.Scan(&v.VideoID, &v.Title,&v.Count); err != nil {
			return nil, err
		}
		v.Link="/play/"+v.VideoID
		ansList = append(ansList, v)
	}
	//fmt.Println(len(ansList))
	return &ansList, nil
}

func (v *videoRepository) Delete(videoID string) error {
	panic("implement me")
}

func NewVideoRepository(saveDir string) repository.VideoRepository{
	if saveDir==""{
		saveDir="./video"
	}
	return &videoRepository{saveDir:saveDir}
}


func (v *videoRepository) FindByVideoID(videoID string)*domain.VideoInfo {
	row:=v.db.QueryRow("SELECT * from video videoID=?",videoID)
	vi:=domain.VideoInfo{
		Title:   "",
		VideoID: "",
		Count:   0,
		Link:    "",
	}
	if err := row.Scan(&vi.Title, &vi.VideoID, &vi.Count); err != nil {
		return nil
	}

	return &vi
}

var _ repository.VideoRepository = &videoRepository{}
