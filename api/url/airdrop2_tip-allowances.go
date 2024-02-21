package degenapiurl

// AirDrop2TipAllowances returns the URL for the following degen API end-point:
//
//	/api/airdrop2/tip-allowances
func AirDrop2TipAllowances() string {
	const request = "/api/airdrop2/tip-allowances"
	const url string = scheme + "://" + request

	return url
}
