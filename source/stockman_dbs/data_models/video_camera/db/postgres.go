package db

import (
	"stockman/source/stockman_dbs/client/postgresql"

	"github.com/sirupsen/logrus"
)

type repository struct {
	client postgresql.Client
	logger *logrus.Logger
}

func NewRepository(client postgresql.Client, logger *logrus.Logger) *repository {
	return &repository{
		client,
		logger,
	}
}

// func (r *repository) Create(ctx context.Context, vc videocamera.CameraCreateDTO) error {
// 	q := `
// 		INSERT INTO video_camera
// 			(address, port, login, password, link)
// 		VALUES
// 			($1, $2, $3, $4, $5)
// 		RETURNING link
// 	`

// }
