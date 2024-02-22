package degenapi

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errNilHTTPResponse = erorr.Error("degen: nil HTTP-response")
	errNilPoints       = erorr.Error("degen: nil points")
	errNilReadCloser   = erorr.Error("degen: nil read-closer")
	errNilTipAllowance = erorr.Error("degen: nil tip-allowance")
)
