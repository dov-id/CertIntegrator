package data

import (
	"errors"
	"io"
	"net/http"
	"time"
)

type Network string

func (n Network) String() string { return string(n) }

var (
	ErrNoContract                  = errors.New("no contract was found")
	ErrFailedToCastKey             = errors.New("failed to cast public key to ECDSA")
	ErrFailedToCastClients         = errors.New("failed to cast interface{} to map[data.Network]*ethclient.Client")
	ErrFailedToCastCertIntegrators = errors.New("failed to cast interface{} to map[data.Network]*contracts.CertIntegratorContract")
	ErrReplacementTxUnderpriced    = errors.New("replacement transaction underpriced")
	ErrWrongSignatureValue         = errors.New("wrong signature value")
	ErrWrongSignatureValues        = errors.New("wrong signature values")
	ErrInvalidPublicKey            = errors.New("invalid public key")
	ErrNoPublicKey                 = errors.New("no public key was found")
	ErrNoSuchKey                   = errors.New("no such key in storage")
	ErrTxWithoutSignature          = errors.New("server returned transaction without signature")
	ErrNotString                   = errors.New("the value is not a string")
	ErrInvalidEthAddress           = errors.New("given value is invalid ethereum address")
)

const (
	MaxMTreeLevel  = 64
	NetworksAmount = 3
)

const (
	EthereumNetwork Network = "ethereum"
	PolygonNetwork  Network = "polygon"
	QNetwork        Network = "q"
)

const (
	NetworkClients          = "network clients"
	CertIntegratorContracts = "feedback registries contracts"
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
