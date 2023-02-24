package stockmanvideomanager

import (
	"context"
	core "stockman/source/stockman_core"
	videocamera "stockman/source/stockman_dbs/data_models/video_camera"
)

func AddNewCamera(ctx context.Context, event *core.Event) {
	newCameraDTO := videocamera.CameraCreateDTO{}
	event.LoadInput(&newCameraDTO)
	CreateNewCamera(ctx, newCameraDTO)
}

func AllCameras() []videocamera.CameraReadDTO {
	return make([]videocamera.CameraReadDTO, 0)
}

func UpdateCamera(videocamera.CameraUpdateDTO) {}

func VideoStream() {}
