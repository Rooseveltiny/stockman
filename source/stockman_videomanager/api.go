package stockmanvideomanager

import videocamera "stockman/source/stockman_dbs/data_models/video_camera"

func AddNewCamera() videocamera.CameraReadDTO {

	return videocamera.CameraReadDTO{}
}

func AllCameras() []videocamera.CameraReadDTO {
	return make([]videocamera.CameraReadDTO, 0)
}

func UpdateCamera(videocamera.CameraUpdateDTO) {}

func VideoStream() {}
