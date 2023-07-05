package data

import (
	"io"
	"net/http"
	"time"
)

const (
	NilCancelFunctionErr        = "ctx cancel function is nil"
	NoContractErr               = "no contract was found"
	FailedToCastKeyErr          = "failed to cast public key to ECDSA"
	FailedToCastIntErr          = "failed to cast interface{} to int64"
	ReplacementTxUnderpricedErr = "replacement transaction underpriced"
	WrongSignatureValueErr      = "wrong signature value"
	WrongSignatureValuesErr     = "wrong signature values"
	InvalidPublicKeyErr         = "invalid public key"
	NoPublicKeyErr              = "no public key was found"
	MaxAttemptsAmountErr        = "attempts amount to get public key reached max value"
	NoSuchKeyErr                = "no such key in storage"
)

const (
	MaxMTreeLevel  = 64
	IndexerTimeout = 30
)

const (
	MetamaskNetwork = "metamask"
	InfuraNetwork   = "infura"
	EthereumNetwork = "ethereum"
	PolygonNetwork  = "polygon"
	QNetwork        = "q"
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
