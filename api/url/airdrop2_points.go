package degenapiurl

import (
	gourl "net/url"
)

// AirDrop2PointsUsingWalletAddress returns the URL for the follow degen API end-point:
//
//	/api/airdrop2/points?address=<wallet>
func AirDrop2PointsUsingWalletAddress(walletAddress string) string {
	var buffer [bufferSize]byte
	var p []byte = buffer[0:0]

	var url string
	{
		const requestPrefix = "/api/airdrop2/points?address="
		const urlPrefix string = scheme + "://" + requestPrefix

		p = append(p, urlPrefix...)
		p = append(p, gourl.QueryEscape(walletAddress)...)

		url = string(p)
	}

	return url
}
