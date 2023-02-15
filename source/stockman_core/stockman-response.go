package core

/*
Response from stockman service. Has err if it occured.
*/
type StockmanResponse[T any] struct {
	Err     error
	DTOData T
}
