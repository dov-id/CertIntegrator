package responses

import (
	"fmt"

	"github.com/dov-id/cert-integrator-svc/resources"
)

func newSMTProof(id int64, root, key, value string, proof []string) resources.SmtProof {
	return resources.SmtProof{
		Key: resources.NewKeyInt64(id, resources.SMT_PROOF),
		Attributes: resources.SmtProofAttributes{
			NodeKey:   fmt.Sprintf("0x%s", key),
			NodeValue: fmt.Sprintf("0x%s", value),
			Proof:     proof,
		},
	}
}

func NewSMTProofResponse(id int64, root, key, value string, proof []string) resources.SmtProofResponse {
	return resources.SmtProofResponse{
		Data: newSMTProof(id, root, key, value, proof),
	}
}
