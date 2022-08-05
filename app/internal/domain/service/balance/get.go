package balance

import (
	"context"

	app_errors "balance-service/app/pkg/errors"
	"github.com/pkg/errors"
)

func (s *balanceService) Get(ctx context.Context, userID int64, abbr string) (string, error) {

	// get balance
	money, err := s.balance.GetBalance(ctx, userID)
	if err != nil {
		if internal, ok := app_errors.ToInternal(err); ok {
			return "", internal.Wrap("get balance")
		}

		return "", app_errors.Wrap(err, "get balance")
	}

	// //

	// exchange rate
	if abbr != "" {
		// get currency
		c, err := s.currency.Get(ctx, abbr)
		if err != nil {
			if internal, ok := app_errors.ToInternal(err); ok {
				return "", internal.Wrap("get currency")
			}
			return "", errors.Wrap(err, "get currency")
		}

		money = money.Div(c) // convert value
	}

	// format '1.00000...' -> '1.00'
	return money.StringFixed(2), nil
}
