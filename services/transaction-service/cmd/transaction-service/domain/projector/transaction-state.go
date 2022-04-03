package projector

type TransactionState struct {
	ValuePaid   int64  `json:"value_paid"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	DeletedAt   int64  `json:"deleted_at,omitempty"`
}
