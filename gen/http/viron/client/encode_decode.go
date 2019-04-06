// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// Viron HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/tonouchi510/goa2-sample/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	viron "github.com/tonouchi510/goa2-sample/gen/viron"
	vironviews "github.com/tonouchi510/goa2-sample/gen/viron/views"
	goahttp "goa.design/goa/http"
)

// BuildAuthtypeRequest instantiates a HTTP request object with method and path
// set to call the "Viron" service "authtype" endpoint
func (c *Client) BuildAuthtypeRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AuthtypeVironPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Viron", "authtype", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeAuthtypeResponse returns a decoder for responses returned by the Viron
// authtype endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeAuthtypeResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body AuthtypeResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Viron", "authtype", err)
			}
			p := NewAuthtypeVironAuthtypeCollectionOK(body)
			view := "default"
			vres := vironviews.VironAuthtypeCollection{p, view}
			if err = vironviews.ValidateVironAuthtypeCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("Viron", "authtype", err)
			}
			res := viron.NewVironAuthtypeCollection(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Viron", "authtype", resp.StatusCode, string(body))
		}
	}
}

// BuildVironMenuRequest instantiates a HTTP request object with method and
// path set to call the "Viron" service "viron_menu" endpoint
func (c *Client) BuildVironMenuRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: VironMenuVironPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Viron", "viron_menu", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeVironMenuResponse returns a decoder for responses returned by the
// Viron viron_menu endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeVironMenuResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body VironMenuResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Viron", "viron_menu", err)
			}
			p := NewVironMenuViewOK(&body)
			view := "default"
			vres := &vironviews.VironMenu{p, view}
			if err = vironviews.ValidateVironMenu(vres); err != nil {
				return nil, goahttp.ErrValidationError("Viron", "viron_menu", err)
			}
			res := viron.NewVironMenu(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Viron", "viron_menu", resp.StatusCode, string(body))
		}
	}
}

// unmarshalVironAPIResponseBodyToVironviewsVironAPIView builds a value of type
// *vironviews.VironAPIView from a value of type *VironAPIResponseBody.
func unmarshalVironAPIResponseBodyToVironviewsVironAPIView(v *VironAPIResponseBody) *vironviews.VironAPIView {
	res := &vironviews.VironAPIView{
		Method: v.Method,
		Path:   v.Path,
	}

	return res
}
