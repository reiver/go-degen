package degenapi

import (
	"sourcecode.social/reiver/go-opt"
)

type AirDrop2DailyTipAllowance struct {
	FID                opt.Optional[string] `json:"fid"`
	SnapshotDate       opt.Optional[string] `json:"snapshot_date"`
	UserRank           opt.Optional[string] `json:"user_rank"`
	WalletAddress      opt.Optional[string] `json:"wallet_address"`
	AvatarURL          opt.Optional[string] `json:"avatar_url"`
	DisplayName        opt.Optional[string] `json:"display_name"`
	ReactionsPerCast   opt.Optional[string] `json:"reactions_per_cast"`
	TipAllowance       opt.Optional[string] `json:"tip_allowance"`
	RemainingAllowance opt.Optional[string] `json:"remaining_allowance"`
}
