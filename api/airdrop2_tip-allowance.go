package degenapi

import (
	"encoding/json"
	"net/http"

	"sourcecode.social/reiver/go-opt"

	"github.com/reiver/go-degen/api/url"
)

// GetAirDrop2TipAllowanceUsingWalletAddress calls the following degen API end-point:
//
//	/api/airdrop2/tip-allowance?address=<wallet>
func GetAirDrop2TipAllowanceUsingWalletAddress(dailyTipAllowance *AirDrop2DailyTipAllowance, walletAddress string) error {
	if nil == dailyTipAllowance {
		return errNilTarget
	}

	var url string = degenapiurl.AirDrop2TipAllowanceUsingWalletAddress(walletAddress)

	resp, err := http.Get(url)
	if nil != err {
		return err
	}
	if nil == resp {
		return errNilHTTPResponse
	}

	return handleHTTPResponse(resp, func(p []byte)error{
		return unmarshalAirDrop2TipAllowance(dailyTipAllowance, p)
	})
}

// GetAirDrop2TipAllowanceUsingFarcasterID calls the following degen API end-point:
//
//	/api/airdrop2/tip-allowance?fid=<Farcaster ID>
func GetAirDrop2TipAllowanceUsingFarcasterID(dailyTipAllowance *AirDrop2DailyTipAllowance,farcasterID string) error {
	if nil == dailyTipAllowance {
		return errNilTarget
	}

	var url string = degenapiurl.AirDrop2TipAllowanceUsingFarcasterID(farcasterID)

	resp, err := http.Get(url)
	if nil != err {
		return err
	}
	if nil == resp {
		return errNilHTTPResponse
	}

	return handleHTTPResponse(resp, func(p []byte)error{
		return unmarshalAirDrop2TipAllowance(dailyTipAllowance, p)
	})
}

func unmarshalAirDrop2TipAllowance(dailyTipAllowance *AirDrop2DailyTipAllowance, p []byte) error {
	if nil == dailyTipAllowance {
		return errNilTarget
	}

	var result []map[string]string = []map[string]string{}

	err := json.Unmarshal(p, &result)
	if nil != err {
		return err
	}

	if len(result) < 1 {
		return nil
	}
	var result0 map[string]string = result[0]

	{
		const name string = "fid"

		value, found := result0[name]
		if found {
			dailyTipAllowance.FID = opt.Something(value)
		}
	}

	{
		const name string = "snapshot_date"

		value, found := result0[name]
		if found {
			dailyTipAllowance.SnapshotDate = opt.Something(value)
		}
	}

	{
		const name string = "user_rank"

		value, found := result0[name]
		if found {
			dailyTipAllowance.UserRank = opt.Something(value)
		}
	}

	{
		const name string = "wallet_address"

		value, found := result0[name]
		if found {
			dailyTipAllowance.WalletAddress = opt.Something(value)
		}
	}

	{
		const name string = "avatar_url"

		value, found := result0[name]
		if found {
			dailyTipAllowance.AvatarURL = opt.Something(value)
		}
	}

	{
		const name string = "display_name"

		value, found := result0[name]
		if found {
			dailyTipAllowance.DisplayName = opt.Something(value)
		}
	}

	{
		const name string = "reactions_per_cast"

		value, found := result0[name]
		if found {
			dailyTipAllowance.ReactionsPerCast = opt.Something(value)
		}
	}

	{
		const name string = "tip_allowance"

		value, found := result0[name]
		if found {
			dailyTipAllowance.TipAllowance = opt.Something(value)
		}
	}

	{
		const name string = "remaining_allowance"

		value, found := result0[name]
		if found {
			dailyTipAllowance.RemainingAllowance = opt.Something(value)
		}
	}

	return nil
}
