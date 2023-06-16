package requests

import (
	"encoding/json"
	"net/http"

	"github.com/dov-id/cert-integrator-svc/resources"
	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type GenerateProofRequest struct {
	Data resources.GenProof `json:"data"`
}

func NewGenerateProofRequest(r *http.Request) (GenerateProofRequest, error) {
	var request GenerateProofRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *GenerateProofRequest) validate() error {
	return validation.Errors{
		"node_key": validation.Validate(&r.Data.Attributes.NodeKey, validation.Required),
		"contract": validation.Validate(&r.Data.Attributes.Contract, validation.Required),
	}.Filter()
}
