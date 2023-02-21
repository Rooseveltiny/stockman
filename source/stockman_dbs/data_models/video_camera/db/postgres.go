package db

import (
	"context"
	"fmt"
	"stockman/source/stockman_dbs/client/postgresql"
	videocamera "stockman/source/stockman_dbs/data_models/video_camera"

	"github.com/sirupsen/logrus"
)

type repository struct {
	client postgresql.Client
	logger *logrus.Logger
}

func Create(ctx context.Context, vc *videocamera.VideoCamera) error {
	q := `
		INSERT INTO video_camera
			(address, port, login, password, link)
		VALUES
			($1, $2, $3, $4, $5)
		RETURNING link
	`

	fmt.Println(q)
	return nil
}
