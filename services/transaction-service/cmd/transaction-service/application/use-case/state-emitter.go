package usecase

import "transaction-service/cmd/transaction-service/domain/projector"

type StateEmitter interface {
	Emit(state projector.TransactionState, id, accountID string)
}
