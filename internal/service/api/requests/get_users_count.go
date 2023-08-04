package requests

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
)

type GetUsersCountRequest struct {
	Course *string `filter:"course"`
}

func NewGetUsersCountRequest(r *http.Request) (GetUsersCountRequest, error) {
	var request GetUsersCountRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, errors.Wrap(err, "failed to decode url")
	}

	return request, request.validate()
}

func (r *GetUsersCountRequest) validate() error {
	course := ""
	if r.Course != nil {
		course = *r.Course
	}

	return validation.Errors{
		"course": validation.Validate(
			&course, validation.Required, validation.By(MustBeValidEthAddress),
		),
	}.Filter()
}
