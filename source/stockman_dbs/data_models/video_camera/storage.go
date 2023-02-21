package videocamera

import "context"

type Repository interface {
	Create(ctx context.Context, vc *VideoCamera) error
}
