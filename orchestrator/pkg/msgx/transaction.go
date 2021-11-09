package msgx

type TransactionInfo struct {
	CorrelationID string `json:"correlationID"`
	Action        string `json:"action"`
}
