package requests

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
)

type GetPublicKeyRequest struct {
	Address *string `filter:"address"`
	Course  *string `filter:"course"`
}

func NewGetPublicKeyRequest(r *http.Request) (GetPublicKeyRequest, error) {
	var request GetPublicKeyRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, errors.Wrap(err, "failed to decode url")
	}

	return request, request.validate()
}

func (r GetPublicKeyRequest) validate() error {
	course := ""
	if r.Course != nil {
		course = *r.Course
	}

	address := ""
	if r.Address != nil {
		address = *r.Address
	}

	return validation.Errors{
		"address": validation.Validate(
			&address, validation.Required, validation.By(MustBeValidEthAddress),
		),
		"course": validation.Validate(
			&course, validation.Required, validation.By(MustBeValidEthAddress),
		),
	}.Filter()
}
