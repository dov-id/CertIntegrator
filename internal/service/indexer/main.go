package indexer

import (
	"context"

	"github.com/dov-id/cert-integrator-svc/contracts"
	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/data/postgres"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Indexer interface {
	Run(ctx context.Context)
}

type indexer struct {
	Type string

	cfg config.Config
	log *logan.Entry

	issuerCh  chan string
	Addresses []string
	Blocks    map[string]int64

	MasterQ data.MasterQ

	Clients         map[data.Network]*ethclient.Client
	CertIntegrators map[data.Network]*contracts.CertIntegratorContract
}

type newIndexerParams struct {
	name            string
	cfg             config.Config
	ctx             context.Context
	issuerCh        chan string
	issuerAddresses []string
	issuerBlocks    map[string]int64
	fabricAddresses []string
	fabricBlocks    map[string]int64
	clients         map[data.Network]*ethclient.Client
	certIntegrators map[data.Network]*contracts.CertIntegratorContract
}

func Run(cfg config.Config, ctx context.Context) {
	params, err := prepareIndexerParams(ctx, cfg)
	if err != nil {
		panic(errors.Wrap(err, "failed to prepare indexer params"))
	}

	params.name = IssuerContract
	NewIndexer(*params).Run(ctx)

	params.name = FabricContract
	NewIndexer(*params).Run(ctx)
}

func NewIndexer(params newIndexerParams) Indexer {
	addresses := params.issuerAddresses
	blocks := params.issuerBlocks
	if params.name == FabricContract {
		addresses = params.fabricAddresses
		blocks = params.fabricBlocks
	}

	return &indexer{
		cfg:             params.cfg,
		log:             params.cfg.Log(),
		issuerCh:        params.issuerCh,
		Addresses:       addresses,
		Blocks:          blocks,
		MasterQ:         postgres.NewMasterQ(params.cfg.DB().Clone()),
		Clients:         params.clients,
		CertIntegrators: params.certIntegrators,
	}
}
