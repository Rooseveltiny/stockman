package stockmanvideomanager

import (
	"context"
	videocamera "stockman/source/stockman_dbs/data_models/video_camera"
	"stockman/source/stockman_dbs/data_models/video_camera/vcpostgres"
	logger "stockman/source/stockman_logger"
)

type VideoCamera struct {
	Address  string
	Port     string
	Login    string
	Password string
	Link     string
}

func (vc *VideoCamera) ConnectionEstablished() bool { return false }
func (vc *VideoCamera) VideoStream()                {}

func CreateNewCamera(ctx context.Context, vc videocamera.CameraCreateDTO) (string, error) {
	repo := vcpostgres.NewRepository(ctx, logger.L)
	vcLink, err := repo.Create(ctx, vc)
	if err != nil {
		logger.L.Errorln(err)
		return "", err
	}
	return vcLink, nil
}

func AllVideoCameras() []*VideoCamera { return nil }
func UpdateVideoCamera() error        { return nil }
