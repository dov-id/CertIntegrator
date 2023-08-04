package requests

import (
	"encoding/json"
	"net/http"

	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/resources"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type GetUsersRequest resources.GetUsersRequest

func NewGetUsersRequest(r *http.Request) (GetUsersRequest, error) {
	var request GetUsersRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *GetUsersRequest) validate() error {
	return validation.Errors{
		"ids": validation.Validate(
			&r.Data.Attributes.Ids, validation.Required, validation.By(MustBeUint),
		),
		"course": validation.Validate(
			&r.Data.Attributes.Course, validation.Required, validation.By(MustBeValidEthAddress),
		),
	}.Filter()
}

func MustBeUint(val interface{}) error {
	arr, ok := val.(*[]int64)
	if !ok {
		return data.ErrInvalidInt64Array
	}

	for _, el := range *arr {
		if el < 0 {
			return data.ErrInvalidIdx
		}
	}

	return nil
}
