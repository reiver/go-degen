package degenapi

import (
	"encoding/json"
	"net/http"

	"sourcecode.social/reiver/go-opt"

	"github.com/reiver/go-degen/api/url"
)

type AirDrop2Points struct {
//	FID           opt.Optional[string] `json:"fid"`
//	WalletAddress opt.Optional[string] `json:"wallet_address"`
//	AvatarURL     opt.Optional[string] `json:"avatar_url"`
	DisplayName   opt.Optional[string] `json:"display_name"`
	Points        opt.Optional[string] `json:"points"`
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

func unmarshalAirDrop2Points(target *AirDrop2Points, p []byte) error {
	if nil == target {
		return errNilTarget
	}

	var list []AirDrop2Points

	if err := json.Unmarshal(p, &list); nil != err {
		return err
	}

	if len(list) < 1 {
		return nil
	}

	*target = list[0]
	return nil
}

