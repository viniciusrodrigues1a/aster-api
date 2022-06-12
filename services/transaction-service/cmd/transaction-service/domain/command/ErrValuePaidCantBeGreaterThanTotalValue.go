package command

import "errors"

var ErrValuePaidCantBeGreaterThanTotalValue = errors.New("value paid can't be greater than total value of transaction")
