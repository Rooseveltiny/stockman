package main

/*
Core entity to start stockman system. Here is going to be loaded stockman service and all
other subsystems i.e. stockman rest API backend or DB client.
*/
type StockmanService struct {
	SystemEventsManager *SystemEventsManager
}

func (ss *StockmanService) StartSystem() {

}

func NewStockmanService() *StockmanService {
	return &StockmanService{}
}
