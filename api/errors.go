package degenapi

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errNilHTTPResponse = erorr.Error("degen: nil HTTP-response")
	errNilReadCloser   = erorr.Error("degen: nil read-closer")
	errNilTipAllowance = erorr.Error("degen: nil tip-allowance")
)
