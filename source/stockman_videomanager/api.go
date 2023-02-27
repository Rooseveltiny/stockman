package stockmanvideomanager

import (
	"context"
	core "stockman/source/stockman_core"
	videocamera "stockman/source/stockman_dbs/data_models/video_camera"
	logger "stockman/source/stockman_logger"

	"github.com/dranikpg/go-dto"
)

func RetrieveCamera(ctx context.Context, event *core.Event) {
	var cameraLink string
	vcreadDTO := videocamera.CameraReadDTO{}
	event.LoadInput(&cameraLink)
	videocamera, err := GetVideoCamera(ctx, cameraLink)
	if err != nil {
		event.SetError(&CantFindCamera)
		event.NotifyOutputChanged()
	}
	dto.Map(&vcreadDTO, videocamera)
	event.SetOutput(vcreadDTO)
	event.NotifyOutputChanged()
}

func AddNewCamera(ctx context.Context, event *core.Event) {
	newCameraDTO := videocamera.CameraCreateDTO{}
	event.LoadInput(&newCameraDTO)
	l, err := CreateNewCamera(ctx, newCameraDTO)
	if err != nil {
		logger.L.Errorln(err)
		event.SetError(&CantCreateCamera)
	}
	event.SetOutput(l)
	event.NotifyOutputChanged()
}

func AllCameras() []videocamera.CameraReadDTO {
	return make([]videocamera.CameraReadDTO, 0)
}

func UpdateCamera(videocamera.CameraUpdateDTO) {}

func VideoStream() {}
