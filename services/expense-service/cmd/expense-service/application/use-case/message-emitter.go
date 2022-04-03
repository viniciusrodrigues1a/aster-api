package usecase

import "expense-service/cmd/expense-service/domain/projector"

type StateEmitter interface {
	Emit(state projector.ExpenseState, id string, accountId string)
}
