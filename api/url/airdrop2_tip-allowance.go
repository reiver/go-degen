package degenapiurl

import (
	gourl "net/url"
)

// AirDrop2TipAllowanceUsingWalletAddress returns the URL for the following degen API end-point:
//
//	/api/airdrop2/tip-allowance?address=<wallet>
func AirDrop2TipAllowanceUsingWalletAddress(walletAddress string) string {
	var buffer [bufferSize]byte
	var p []byte = buffer[0:0]

	var url string
	{
		const requestPrefix = "/api/airdrop2/tip-allowance?address="
		const urlPrefix string = scheme + "://" + host + "/" + requestPrefix

		p = append(p, urlPrefix...)
		p = append(p, gourl.QueryEscape(walletAddress)...)

		url = string(p)
	}

	return url
}


// AirDrop2TipPointsUsingFarcasterID returns the URL for the following degen API end-point:
//
//	/api/airdrop2/tip-allowance?fid=<Farcaster ID>
func AirDrop2TipAllowanceUsingFarcasterID(farcasterID string) string {
	var buffer [bufferSize]byte
	var p []byte = buffer[0:0]

	var url string
	{
		const requestPrefix = "/api/airdrop2/tip-allowance?fid="
		const urlPrefix string = scheme + "://" + requestPrefix

		p = append(p, urlPrefix...)
		p = append(p, gourl.QueryEscape(farcasterID)...)

		url = string(p)
	}

	return url
}
