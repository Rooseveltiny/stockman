package stockmanvideomanager

import (
	videocamera "stockman/source/stockman_dbs/data_models/video_camera"
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

func CreateNewCamera(camera videocamera.CameraCreateDTO) *VideoCamera {
	return &VideoCamera{}
}

func AllVideoCameras() []*VideoCamera { return nil }
func UpdateVideoCamera() error        { return nil }
