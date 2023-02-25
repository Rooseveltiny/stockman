package stockmanrestapiserver

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/* handle funcs */
func CreateNewVideoCamera(w http.ResponseWriter, r *http.Request, p httprouter.Params) {}

/* handlers */
var CreateNewVideoCameraHandler *Hand = NewHand(http.MethodPost, "/create_video_camera", CreateNewVideoCamera)

/* routes collection */
var VideocameraHandlers *RoutesCollection = NewRoutesCollection()

func init() {
	VideocameraHandlers.AppendHandle(*CreateNewVideoCameraHandler)
}
