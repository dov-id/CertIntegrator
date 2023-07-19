package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/urlval"
)

type GetPublicKeysRequest struct {
	Course *string `filter:"course"`
	Signer *string `filter:"address"`
}

func NewGetPublicKeysRequest(r *http.Request) (GetPublicKeysRequest, error) {
	var request GetPublicKeysRequest

	err := urlval.Decode(r.URL.Query(), &request)

	return request, err
}
