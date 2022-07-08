package dto

import (
	"fmt"

	"balance-service/app/internal/controller/http/user/dto"
)

type UnblockIN struct {
	UserID int64 `json:"user_id"`
}
type UnblockOUT dto.Status

func (g UnblockIN) Validate() error {
	if g.UserID <= 0 {
		return fmt.Errorf("invalid user id: got %d, expected > 0", g.UserID)
	}
	return nil
}

// //

type BlockIN struct {
	UnblockIN
	Reason string `json:"reason"`
}

func (b BlockIN) Validate() error {
	if b.Reason == "" {
		return fmt.Errorf("invalid reason: got '%s', expected not empty", b.Reason)
	}
	return b.UnblockIN.Validate()
}

type BlockOUT dto.Status
