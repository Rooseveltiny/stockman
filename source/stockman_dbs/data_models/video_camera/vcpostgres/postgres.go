package vcpostgres

import (
	"context"
	"stockman/source/stockman_dbs/client/postgresql"
	videocamera "stockman/source/stockman_dbs/data_models/video_camera"
	logger "stockman/source/stockman_logger"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/sirupsen/logrus"
)

type repository struct {
	client postgresql.Client
	logger *logrus.Logger
}

func NewRepository(ctx context.Context, logger *logrus.Logger) *repository {
	client, _ := postgresql.GetPostgresClient(ctx)
	return &repository{
		client,
		logger,
	}
}

func (r *repository) Create(ctx context.Context, vc videocamera.CameraCreateDTO) (string, error) {
	q :=
		`
		INSERT INTO video_camera
			(address, port, login, password)
		VALUES
			($1, $2, $3, $4)
		RETURNING link
		`
	row := r.client.QueryRow(ctx, q, vc.Address, vc.Port, vc.Login, vc.Password)
	var link string
	err := row.Scan(&link)
	if err != nil {
		r.logger.Errorln(err)
		return "", err
	}
	return link, nil
}

func (r *repository) GetByLink(ctx context.Context, link string) (*videocamera.CameraReadDTO, error) {
	q := `
		SELECT * FROM video_camera
		WHERE link = $1
	`
	var ccd videocamera.CameraReadDTO
	err := pgxscan.Get(ctx, r.client, &ccd, q, link)
	if err != nil {
		logger.L.Errorln(err)
	}
	return &ccd, nil
}
