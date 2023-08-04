package helpers

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	pkgErrors "github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type ScannerResponse struct {
	Status  string               `json:"status"`
	Message string               `json:"message"`
	Result  []ScannerTransaction `json:"result"`
}

type ScannerTransaction struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	TransactionIndex  string `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	IsError           string `json:"isError"`
	TxReceiptStatus   string `json:"txreceipt_status"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	Confirmations     string `json:"confirmations"`
}

type ProcessPubKeyParams struct {
	Ctx        context.Context
	Cfg        config.Config
	Address    common.Address
	UsersQ     data.Users
	ContractId uint64
	Clients    map[data.Network]*ethclient.Client
}

func ProcessPublicKey(params ProcessPubKeyParams) error {
	user, err := params.UsersQ.FilterByAddresses(params.Address.Hex()).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get user")
	}

	if user != nil {
		return nil
	}

	publicKey, err := RetrievePublicKey(params.Ctx, params.Address, params.Cfg.Networks().Networks, params.Clients)
	if err != nil {
		return errors.Wrap(err, "failed to retrieve public key")
	}

	if publicKey == nil {
		return data.ErrNoPublicKey
	}

	err = params.UsersQ.Upsert(data.User{
		Address:    params.Address.Hex(),
		ContractId: params.ContractId,
		PublicKey:  fmt.Sprintf("0x%s", hex.EncodeToString(publicKey)),
	})
	if err != nil {
		return errors.Wrap(err, "failed to upsert user")
	}

	return nil
}

func RetrievePublicKey(
	ctx context.Context,
	address common.Address,
	networks map[data.Network]config.Network,
	clients map[data.Network]*ethclient.Client,
) ([]byte, error) {
	for network, params := range networks {
		requestParams := data.RequestParams{
			Method: http.MethodGet,
			Link:   params.BlockExplorerApiUrl,
			Body:   nil,
			Query: map[string]string{
				"module":  "account",
				"action":  "txlist",
				"address": address.Hex(),
				"apikey":  params.BlockExplorerApiKey,
			},
			Header:  nil,
			Timeout: 100 * time.Second,
		}

		response, err := MakeHttpRequest(ctx, requestParams)
		if err != nil {
			return nil, errors.Wrap(err, "failed to make http request")
		}

		var body ScannerResponse
		if err = json.NewDecoder(response.Body).Decode(&body); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal body")
		}

		if len(body.Result) == 0 {
			continue
		}

		publicKey, err := getPublicKey(ctx, address, body.Result, clients[network])
		if err != nil {
			return nil, errors.Wrap(err, "failed to get public key")
		}

		return publicKey, nil
	}

	return nil, nil
}

func getPublicKey(ctx context.Context, address common.Address, txs []ScannerTransaction, client *ethclient.Client) ([]byte, error) {
	for _, tx := range txs {
		if strings.ToLower(tx.From) != strings.ToLower(address.Hex()) {
			continue
		}

		transaction, _, err := client.TransactionByHash(ctx, common.HexToHash(tx.Hash))
		if pkgErrors.Is(err, data.ErrTxWithoutSignature) {
			continue
		}

		if err != nil {
			return nil, errors.Wrap(err, "failed to get tx by hash")
		}

		chainID, err := client.ChainID(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get chain id")
		}

		publicKey, err := recoverPubKeyFromTx(transaction, types.NewLondonSigner(chainID))
		if err != nil {
			return nil, errors.Wrap(err, "failed to recover public key from transaction")
		}

		return publicKey, nil
	}

	return nil, nil
}

func recoverPubKeyFromTx(transaction *types.Transaction, signer types.Signer) ([]byte, error) {
	vBig, r, s := transaction.RawSignatureValues()

	if vBig.BitLen() > 8 {
		return nil, data.ErrWrongSignatureValue
	}

	v := byte(vBig.Uint64())

	if !crypto.ValidateSignatureValues(v, r, s, true) {
		return nil, data.ErrWrongSignatureValues
	}

	R, S := r.Bytes(), s.Bytes()

	signature := make([]byte, crypto.SignatureLength)
	copy(signature[32-len(R):32], R)
	copy(signature[64-len(S):64], S)
	signature[64] = v

	pubKey, err := crypto.Ecrecover(signer.Hash(transaction).Bytes(), signature)
	if err != nil {
		return nil, errors.Wrap(err, "failed to recover signature")
	}

	if len(pubKey) == 0 || pubKey[0] != 4 {
		return nil, data.ErrInvalidPublicKey
	}

	return pubKey, nil
}
