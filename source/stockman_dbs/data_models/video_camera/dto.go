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
	Address  string `json:"address"`
	Port     string `json:"port"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Link     string `json:"link"`
}
