package core

/*
Core entity to start stockman system. Here is going to be loaded stockman service and all
other subsystems i.e. stockman rest API backend or DB client.
*/

var SystemEvents_Manager *SystemEventsManager

type StockmanService struct {
	SystemEventsManager *SystemEventsManager
}

func (ss *StockmanService) StartSystem() {
	sysevman := NewSystemEventsManager()
	sysevman.RunEventLoop()
	ss.SystemEventsManager = sysevman
	SystemEvents_Manager = sysevman
}

func NewStockmanService() *StockmanService {
	return &StockmanService{}
}

func InitAndRunStockmanService() {
	ss := NewStockmanService()
	ss.StartSystem()
}
