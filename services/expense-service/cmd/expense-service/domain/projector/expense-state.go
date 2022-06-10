package projector

type ExpenseState struct {
	ProductID   *string
	Title       string `json:"title"`
	Description string `json:"description"`
	Value       int64  `json:"value"`
	CreatedAt   int64  `json:"created_at"`
	DeletedAt   int64  `json:"deleted_at,omitempty"`
}
