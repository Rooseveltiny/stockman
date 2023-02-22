package videocamera

import "context"

type Repository interface {
	Create(ctx context.Context, vc CameraCreateDTO) (string, error)
	GetByLink(ctx context.Context, link string) (*CameraReadDTO, error)
}
