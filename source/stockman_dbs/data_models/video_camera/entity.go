package videocamera

import (
	"context"
)

type VideoCamera struct {
	Address  string
	Port     string
	Login    string
	Password string
	Link     string
}

func (vc *VideoCamera) CheckConnectionStatus() {}

func (vc *VideoCamera) VideoStreaming() {}

func (vc *VideoCamera) CreateVideoCamera(ctx context.Context, camera *VideoCamera) error {
	return nil
}

func (vc *VideoCamera) FetchAllVideoCameras() {}

func (vc *VideoCamera) UpdateVideoCamera() {}
