package service

import (
	"context"

	"balance-service/app/pkg/errors"
	"github.com/shopspring/decimal"
)

func (s *balanceService) Get(ctx context.Context, userID int64, abbr string) (string, error) {
	balance, err := s.balance.GetBalance(ctx, userID)
	if err != nil {
		if internal, ok := errors.ToInternal(err); ok {
			return "", internal.Wrap("get balance")
		}

		return "", errors.Wrap(err, "get balance")
	}

	// //
	money := decimal.NewFromInt(balance).Shift(-2)

	// exchange rate
	if abbr != "" {
		// get currency
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

// //

func (s *balanceService) BlockBalance(ctx context.Context, userID int64, reason string) error {
	return s.balance.BlockBalance(ctx, userID, reason)
}

func (s *balanceService) UnblockBalance(ctx context.Context, userID int64) error {
	return s.balance.UnblockBalance(ctx, userID)
}
