package requests

import (
	"encoding/json"
	"net/http"

	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/resources"
	"github.com/ethereum/go-ethereum/common"
	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type GenerateProofRequest resources.GenProofRequest

func NewGenerateProofRequest(r *http.Request) (GenerateProofRequest, error) {
	var request GenerateProofRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *GenerateProofRequest) validate() error {
	return validation.Errors{
		"node_key": validation.Validate(&r.Data.Attributes.NodeKey, validation.Required, validation.By(MustBeValidEthAddress)), //is user address
		"contract": validation.Validate(&r.Data.Attributes.Contract, validation.Required, validation.By(MustBeValidEthAddress)),
	}.Filter()
}

func MustBeValidEthAddress(src interface{}) error {
	raw, ok := src.(string)
	if !ok {
		return data.ErrNotString
	}
	if !common.IsHexAddress(raw) {
		return data.ErrInvalidEthAddress
	}

	return nil
}
