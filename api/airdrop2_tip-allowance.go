package degenapi

import (
	"encoding/json"
	"net/http"

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

func unmarshalAirDrop2TipAllowance(target *AirDrop2DailyTipAllowance, p []byte) error {
	if nil == target {
		return errNilTarget
	}

	var list []AirDrop2DailyTipAllowance

	if err := json.Unmarshal(p, &list); nil != err {
		return err
	}

	if len(list) < 1 {
		return nil
	}

	*target = list[0]
	return nil
}
