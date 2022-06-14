package usecase

import "product-service/cmd/product-service/domain/projector"

type StateEmitter interface {
	Emit(state projector.ProductState, id, accountID string)
}
