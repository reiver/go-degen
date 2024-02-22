package degenapi

import (
	"net/http"
	"io"

	"sourcecode.social/reiver/go-errhttp"
)

func handleHTTPResponse(resp *http.Response, fn func([]byte)error) error {
	if http.StatusOK != resp.StatusCode {
		return errhttp.Return(resp.StatusCode)
	}

        var p []byte
        {
                var readcloser io.ReadCloser = resp.Body
                if nil == readcloser {
                        return errNilReadCloser
                }
                defer readcloser.Close()

                var err error

                p, err = io.ReadAll(readcloser)
                if nil != err {
                        return err
                }
                readcloser.Close()
        }

	return fn(p)
}
