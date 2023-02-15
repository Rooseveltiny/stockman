package core

/*
Response from stockman service. Has err if it occured.
*/
type StockmanResponse[T any] struct {
	Err     error
	DTOData T
}

func NewStockmanResponse[T any](dto T, err error) *StockmanResponse[T] {
	return &StockmanResponse[T]{Err: err, DTOData: dto}
}
