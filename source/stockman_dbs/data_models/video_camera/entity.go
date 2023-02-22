package videocamera

type VideoCamera struct {
	Address  string
	Port     string
	Login    string
	Password string
	Link     string
}

func (vc *VideoCamera) CheckConnectionStatus() {}

func (vc *VideoCamera) VideoStreaming() {}

// func (vc *VideoCamera) CreateVideoCamera(ctx context.Context) (string, error) {
// 	// c := postgresql.NewClient(ctx, )
// 	// repo := db.NewRepository()
// }

func (vc *VideoCamera) FetchAllVideoCameras() {}

func (vc *VideoCamera) UpdateVideoCamera() {}
