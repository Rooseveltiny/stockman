package videocamera

type CameraCreateDTO struct {
	Address  string
	Port     string
	Login    string
	Password string
	Link     string
}

type CameraUpdateDTO struct {
	Address  string
	Port     string
	Login    string
	Password string
	Link     string
}

type CameraReadDTO struct {
	Address  string
	Port     string
	Login    string
	Password string
	Link     string
}
