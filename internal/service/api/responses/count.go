package responses

import "github.com/dov-id/cert-integrator-svc/resources"

func NewCountResponse(amount int64) resources.CountResponse {
	return resources.CountResponse{
		Data: newCount(amount),
	}
}

func newCount(amount int64) resources.Count {
	return resources.Count{
		Key: resources.NewKeyInt64(amount, resources.COUNT),
		Attributes: resources.CountAttributes{
			Amount: amount,
		},
	}
}
