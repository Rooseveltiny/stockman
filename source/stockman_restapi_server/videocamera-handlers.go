package stockmanrestapiserver

import (
	"net/http"
	stockmanapi "stockman/source/stockman_api"

	"encoding/json"

	"github.com/julienschmidt/httprouter"
)

/* handle funcs */
func CreateNewVideoCamera(w http.ResponseWriter, r *http.Request, p httprouter.Params) {}

func GetVideoCameraByUUID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cameraLink := p.ByName("camera_link")
	stockmanResp := stockmanapi.GetVideoCameraDTO(cameraLink)
	if stockmanResp.Err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(stockmanResp.Err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(stockmanResp.DTOData)
	w.Write(data)
}

/* handlers */
var CreateNewVideoCameraHandler *Hand = NewHand(http.MethodPost, "/create_camera", CreateNewVideoCamera, nil)
var GetNewVideoCameraHandler *Hand = NewHand(http.MethodPost, "/get_video_camera/:camera_link", GetVideoCameraByUUID, nil)

/* collection of handlers */
var VideoCameraHandlers_V1 *RoutesCollection = NewRoutesCollection("/api/v1/video_camera")

func init() {
	/* register all handlers */
	VideoCameraHandlers_V1.AppendHandle(*CreateNewVideoCameraHandler)
	VideoCameraHandlers_V1.AppendHandle(*GetNewVideoCameraHandler)
}
