package stockmanapi

import (
	core "stockman/source/stockman_core"
	videocamera "stockman/source/stockman_dbs/data_models/video_camera"
	stockmanvideomanager "stockman/source/stockman_videomanager"
)

/* Video manager API */

/* it is rather flexible to make requests to system. otherwise it can be generic used
func AddVideoCamera(dto videocamera.CameraCreateDTO) *core.StockmanResponse[string] {
	ev := core.NewEvent(stockmanvideomanager.AddNewCamera)
	ev.SetInput(dto)
	core.SystemEvents_Manager.AppendEvent(ev)
	<-ev.OnOutputChanged()
	var d string
	ev.LoadOutput(&d)
	stockmanResponse := core.NewStockmanResponse(d, ev.Error())
	return stockmanResponse
}
*/

/* adding new camera into system with some params */
func AddVideoCamera(dto videocamera.CameraCreateDTO) *core.StockmanResponse[string] {
	return core.ApiRequestShortcutEasyResponse[videocamera.CameraCreateDTO, string](stockmanvideomanager.AddNewCamera, dto)
}

/* getting camera from system by link :uuid */
func GetVideoCameraDTO(link string) *core.StockmanResponse[videocamera.CameraReadDTO] {
	return core.ApiRequestShortcutEasyResponse[string, videocamera.CameraReadDTO](stockmanvideomanager.RetrieveCamera, link)
}
