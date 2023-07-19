package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/urlval"
)

type GetPublicKeyRequest struct {
	Address *string `filter:"address"`
}

func NewGetPublicKeyRequest(r *http.Request) (GetPublicKeyRequest, error) {
	var request GetPublicKeyRequest

	err := urlval.Decode(r.URL.Query(), &request)

	return request, err
}
