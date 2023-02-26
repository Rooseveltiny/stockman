package stockmanrestapiserver

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/* handle funcs */
func CreateNewVideoCamera(w http.ResponseWriter, r *http.Request, p httprouter.Params) {}

/* handlers */
var CreateNewVideoCameraHandler *Hand = NewHand(http.MethodPost, "/create_camera", CreateNewVideoCamera, nil)
var GetNewVideoCameraHandler *Hand = NewHand(http.MethodPost, "/get_video_camera/:camera_link", CreateNewVideoCamera, nil)

func init() {
	/* register all handlers */
	AllHandlers.AppendHandle(*CreateNewVideoCameraHandler)
	AllHandlers.AppendHandle(*GetNewVideoCameraHandler)
}
