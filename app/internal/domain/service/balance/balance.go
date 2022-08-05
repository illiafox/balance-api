package balance

import (
	"context"
)

func (s *balanceService) Change(ctx context.Context, userID int64, amount int64, desc string) error {
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
