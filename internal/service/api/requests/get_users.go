package requests

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
)

type GetUsersRequest struct {
	Course *string `filter:"course"`
	Number *int64  `filter:"number"`
}

func NewGetUsersRequest(r *http.Request) (GetUsersRequest, error) {
	var request GetUsersRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, errors.Wrap(err, "failed to decode url")
	}

	return request, request.validate()
}

func (r *GetUsersRequest) validate() error {
	course := ""
	if r.Course != nil {
		course = *r.Course
	}

	number := int64(0)
	if r.Number != nil {
		number = *r.Number
	}

	return validation.Errors{
		"number": validation.Validate(
			&number, validation.Required,
		),
		"course": validation.Validate(
			&course, validation.Required, validation.By(MustBeValidEthAddress),
		),
	}.Filter()
}
