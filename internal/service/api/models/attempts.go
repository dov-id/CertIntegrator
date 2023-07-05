package models

import "github.com/dov-id/cert-integrator-svc/resources"

func NewAttemptsResponse(amount int64) resources.AttemptsResponse {
	return resources.AttemptsResponse{
		Data: newAttempts(amount),
	}
}

func newAttempts(amount int64) resources.Attempts {
	return resources.Attempts{
		Key: resources.NewKeyInt64(amount, resources.ATTEMPTS),
		Attributes: resources.AttemptsAttributes{
			AttemptsLeft: amount,
		},
	}
}
