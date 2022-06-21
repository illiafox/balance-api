package service

import (
	"context"

	"balance-service/app/pkg/errors"
	"github.com/shopspring/decimal"
)

// stores 100
var divisor = decimal.NewFromInt(100)

func (s *balanceService) Get(ctx context.Context, userID int64, abbr string) (string, error) {
	// GetBalance Balance
	balance, err := s.balance.GetBalance(ctx, userID)
	if err != nil {
		if internal, ok := errors.ToInternal(err); ok {
			return "", internal.Wrap("get balance")
		}

		return "", errors.Wrap(err, "get balance")
	}

	money := decimal.NewFromInt(balance).Div(divisor)

	// GetBalance exchange rate
	if abbr != "" {
		c, err := s.currency.Get(ctx, abbr)
		if err != nil {
			if internal, ok := errors.ToInternal(err); ok {
				return "", internal.Wrap("get currency")
			}

			return "", errors.Wrap(err, "get currency")
		}

		money = money.Div(c)
	}

	// format 100 -> '1.00'
	return money.StringFixed(2), nil
}

func (s *balanceService) Change(ctx context.Context, userID, amount int64, desc string) error {
	return s.balance.ChangeBalance(ctx, userID, amount, desc)
}
func (s *balanceService) Transfer(ctx context.Context, fromID, toID, amount int64, desc string) error {
	return s.balance.Transfer(ctx, fromID, toID, amount, desc)
}
func (s *balanceService) ChangeOwner(ctx context.Context, oldUserID int64, newUserID int64) error {
	return s.balance.ChangeOwner(ctx, oldUserID, newUserID)
}
