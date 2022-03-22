package money

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrCurrencyNotEqual = errors.New("currency not equal")

type Money struct {
	amount   int64
	currency *Currency
}

func New(amount int64, currency *Currency) *Money {
	return &Money{
		amount:   amount,
		currency: currency,
	}
}

func (m *Money) Add(target *Money) (*Money, error) {
	return operate(m, target, func(a, b int64) int64 { return a + b })
}

func (m *Money) Subtract(target *Money) (*Money, error) {
	return operate(m, target, func(a, b int64) int64 { return a - b })
}

func (m *Money) Multiply(target *Money) (*Money, error) {
	return operate(m, target, func(a, b int64) int64 { return a * b })
}

func (m *Money) Divide(target *Money) (*Money, error) {
	return operate(m, target, func(a, b int64) int64 { return a / b })
}

func operate(m1 *Money, m2 *Money, fn func(int64, int64) int64) (*Money, error) {
	if !m1.currency.IsEqualTo(m2.currency) {
		return nil, ErrCurrencyNotEqual
	}

	return New(fn(m1.amount, m2.amount), m1.currency), nil
}

func (m *Money) String() string {
	padChar := "0"
	padSize := m.currency.DecimalDigits + 1

	padStr := strconv.FormatInt(m.amount/100, 10)
	paddedStr := padStr + "." + strings.Repeat(padChar, padSize-len(padStr))

	return fmt.Sprintf("%s %s", m.currency.Symbol, paddedStr)
}
