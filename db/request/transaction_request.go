package request

type TransactionDetail struct {
	ID        int     `json:"id"`
	Customer  string  `json:"customer"`
	Price     float32 `json:"price"`
	Quantity  int     `json:"quantity"`
	Timestamp string  `json:"timestamp"`
}
type TransactionRequest struct {
	RequestId   int                 `json:"request_id"`
	Transaction []TransactionDetail `json:"transaction"`
}
