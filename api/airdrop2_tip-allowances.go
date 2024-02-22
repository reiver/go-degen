package degenapi

import (
	"encoding/json"
	"net/http"

	"github.com/reiver/go-degen/api/url"
)

// GetAirDrop2TipAllowances calls the following degen API end-point:
//
//	/api/airdrop2/tip-allowances
func GetAirDrop2TipAllowances(dailyTipAllowances *[]AirDrop2DailyTipAllowance) error {
	if nil == dailyTipAllowances {
		return errNilTarget
	}

	var url string = degenapiurl.AirDrop2TipAllowances()

	resp, err := http.Get(url)
	if nil != err {
		return err
	}
	if nil == resp {
		return errNilHTTPResponse
	}

	return handleHTTPResponse(resp, func(p []byte)error{
		return unmarshalAirDrop2TipAllowances(dailyTipAllowances, p)
	})
}

func unmarshalAirDrop2TipAllowances(target *[]AirDrop2DailyTipAllowance, p []byte) error {
	if nil == target {
		return errNilTarget
	}

	return json.Unmarshal(p, &target)
}
