package stockmanapi

import (
	core "stockman/source/stockman_core"
	videocamera "stockman/source/stockman_dbs/data_models/video_camera"
	stockmanvideomanager "stockman/source/stockman_videomanager"
)

/* Video manager API */
func AddVideoCamera(dto videocamera.CameraCreateDTO) *core.StockmanResponse[string] {
	ev := core.NewEvent(stockmanvideomanager.AddNewCamera)
	ev.SetInput(dto) /* setting new camera dto */
	core.SystemEvents_Manager.AppendEvent(ev)
	<-ev.OnOutputChanged()
	var d string
	ev.LoadOutput(&d)
	stockmanResponse := core.NewStockmanResponse(d, ev.Error())
	return stockmanResponse
}

func GetVideoCameraDTO(link string) *core.StockmanResponse[videocamera.CameraReadDTO] {
	ev := core.NewEvent(stockmanvideomanager.RetrieveCamera)
	ev.SetInput(link)
	core.SystemEvents_Manager.AppendEvent(ev)
	<-ev.OnOutputChanged()
	var data videocamera.CameraReadDTO
	ev.LoadOutput(&data)
	r := core.NewStockmanResponse(data, ev.Error())
	return r
}
