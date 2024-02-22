package degenapi

import (
	"encoding/json"
	"net/http"

	"sourcecode.social/reiver/go-opt"

	"github.com/reiver/go-degen/api/url"
)

type AirDrop2Points struct {
//	FID           opt.Optional[string]
//	WalletAddress opt.Optional[string]
//	AvatarURL     opt.Optional[string]
	DisplayName   opt.Optional[string]
	Points        opt.Optional[string]
}

// GetAirDrop2PointsUsingWalletAddress calls the following degen API end-point:
//
//	/api/airdrop2/points?address=<wallet>
func GetAirDrop2PointsUsingWalletAddress(points *AirDrop2Points, walletAddress string) error {
	if nil == points {
		return errNilTarget
	}

	var url string = degenapiurl.AirDrop2PointsUsingWalletAddress(walletAddress)

	resp, err := http.Get(url)
	if nil != err {
		return err
	}
	if nil == resp {
		return errNilHTTPResponse
	}

	return handleHTTPResponse(resp, func(p []byte)error{
		return unmarshalAirDrop2Points(points, p)
	})
}

func unmarshalAirDrop2Points(points *AirDrop2Points, p []byte) error {
	if nil == points {
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

//	{
//		const name string = "fid"
//
//		value, found := result0[name]
//		if found {
//			points.FID = opt.Something(value)
//		}
//	}

//	{
//		const name string = "wallet_address"
//
//		value, found := result0[name]
//		if found {
//			points.WalletAddress = opt.Something(value)
//		}
//	}

//	{
//		const name string = "avatar_url"
//
//		value, found := result0[name]
//		if found {
//			points.AvatarURL = opt.Something(value)
//		}
//	}

	{
		const name string = "display_name"

		value, found := result0[name]
		if found {
			points.DisplayName = opt.Something(value)
		}
	}

	return nil
}
