// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// Viron HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/tonouchi510/goa2-sample/design

package server

import (
	"context"
	"net/http"

	vironviews "github.com/tonouchi510/goa2-sample/gen/viron/views"
	goahttp "goa.design/goa/http"
)

// EncodeAuthtypeResponse returns an encoder for responses returned by the
// Viron authtype endpoint.
func EncodeAuthtypeResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(vironviews.VironAuthtypeCollection)
		enc := encoder(ctx, w)
		body := NewVironAuthtypeResponseCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeVironMenuResponse returns an encoder for responses returned by the
// Viron viron_menu endpoint.
func EncodeVironMenuResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*vironviews.VironMenu)
		enc := encoder(ctx, w)
		body := NewVironMenuResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// marshalVironviewsVironAPIViewToVironAPIResponseBody builds a value of type
// *VironAPIResponseBody from a value of type *vironviews.VironAPIView.
func marshalVironviewsVironAPIViewToVironAPIResponseBody(v *vironviews.VironAPIView) *VironAPIResponseBody {
	res := &VironAPIResponseBody{
		Method: *v.Method,
		Path:   *v.Path,
	}

	return res
}
