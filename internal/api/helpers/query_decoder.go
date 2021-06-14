package apihelpers

import (
	// External
	"net/url"

	"github.com/gorilla/schema"
	// Internal
)

func createQueryParamsDecoder() *schema.Decoder {
	var decoder = schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	return decoder
}

func DecodeQueryStringParams(url *url.URL, resultPtr interface{}) error {
	var decoder = createQueryParamsDecoder()
	return decoder.Decode(resultPtr, url.Query())
}
