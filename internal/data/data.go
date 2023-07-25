package data

import (
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/dov-id/cert-integrator-svc/internal/types"
)

var (
	ErrNoContract               = errors.New("no contract was found")
	ErrFailedToCastKey          = errors.New("failed to cast public key to ECDSA")
	ErrFailedToCastInt          = errors.New("failed to cast interface{} to int64")
	ErrReplacementTxUnderpriced = errors.New("replacement transaction underpriced")
	ErrWrongSignatureValue      = errors.New("wrong signature value")
	ErrWrongSignatureValues     = errors.New("wrong signature values")
	ErrInvalidPublicKey         = errors.New("invalid public key")
	ErrNoPublicKey              = errors.New("no public key was found")
	ErrMaxAttemptsAmount        = errors.New("attempts amount to get public key reached max value")
	ErrNoSuchKey                = errors.New("no such key in storage")
	ErrTxWithoutSignature       = errors.New("server returned transaction without signature")
	ErrNotString                = errors.New("the value is not a string")
	ErrInvalidEthAddress        = errors.New("given value is invalid ethereum address")
)

const (
	MaxMTreeLevel = 64
)

const (
	EthereumNetwork types.Network = "ethereum"
	PolygonNetwork  types.Network = "polygon"
	QNetwork        types.Network = "q"
)

type RequestParams struct {
	Method  string
	Link    string
	Body    []byte
	Query   map[string]string
	Header  map[string]string
	Timeout time.Duration
}

type ResponseParams struct {
	Body       io.ReadCloser
	Header     http.Header
	StatusCode int
}
