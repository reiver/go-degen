package degenapi

import (
	"sourcecode.social/reiver/go-opt"
)

type AirDrop2DailyTipAllowance struct {
	FID                opt.Optional[string]
	SnapshotDate       opt.Optional[string]
	UserRank           opt.Optional[string]
	WalletAddress      opt.Optional[string]
	AvatarURL          opt.Optional[string]
	DisplayName        opt.Optional[string]
	ReactionsPerCast   opt.Optional[string]
	TipAllowance       opt.Optional[string]
	RemainingAllowance opt.Optional[string]
}
