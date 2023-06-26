package models

import "github.com/dov-id/cert-integrator-svc/resources"

func NewPublicKeyResponse(address, publicKey string) resources.PublicKResponse {
	return resources.PublicKResponse{
		Data: newPublicKey(address, publicKey),
	}
}

func newPublicKey(address, publicKey string) resources.PublicK {
	return resources.PublicK{
		Key: resources.Key{
			ID:   address,
			Type: resources.PUBLIC_KEY,
		},
		Attributes: resources.PublicKAttributes{
			PublicKey: publicKey,
		},
	}
}
