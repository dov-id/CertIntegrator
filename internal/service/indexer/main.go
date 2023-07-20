package indexer

import (
	"context"
	"sync"

	"github.com/dov-id/cert-integrator-svc/contracts"
	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/data/postgres"
	"github.com/dov-id/cert-integrator-svc/internal/service/storage"
	"github.com/dov-id/cert-integrator-svc/internal/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Indexer interface {
	Run(ctx context.Context)
}

type indexer struct {
	cfg config.Config
	log *logan.Entry

	issuerCh  chan string
	Addresses []string
	Blocks    []int64

	ContractsQ data.Contracts
	UsersQ     data.Users

	Cancel context.CancelFunc
	wg     *sync.WaitGroup

	Clients         map[types.Network]*ethclient.Client
	CertIntegrators map[types.Network]*contracts.CertIntegratorContract

	dailyStorage storage.DailyStorage
}

type newIndexerParams struct {
	cfg             config.Config
	ctx             context.Context
	issuerCh        chan string
	issuerAddresses []string
	issuerBlocks    []int64
	fabricAddresses []string
	fabricBlocks    []int64
	cancel          context.CancelFunc
	wg              *sync.WaitGroup
	clients         map[types.Network]*ethclient.Client
	certIntegrators map[types.Network]*contracts.CertIntegratorContract
}

func Run(cfg config.Config, ctx context.Context) {
	cancelCtx, cancelFn := context.WithCancel(ctx)
	var wg sync.WaitGroup

	params, err := prepareIndexerParams(cfg)
	if err != nil {
		panic(errors.Wrap(err, "failed to prepare indexer params"))
	}

	wg.Add(1)

	params.ctx = cancelCtx
	params.wg = &wg
	params.cancel = nil

	NewIndexer(*params).Run(cancelCtx)

	params.cancel = cancelFn

	NewIndexer(*params).Run(ctx)
}

func NewIndexer(params newIndexerParams) Indexer {
	addresses := params.issuerAddresses
	blocks := params.issuerBlocks
	if params.cancel != nil {
		addresses = params.fabricAddresses
		blocks = params.fabricBlocks
	}

	return &indexer{
		cfg:             params.cfg,
		log:             params.cfg.Log(),
		issuerCh:        params.issuerCh,
		Addresses:       addresses,
		Blocks:          blocks,
		ContractsQ:      postgres.NewContractsQ(params.cfg.DB().Clone()),
		UsersQ:          postgres.NewUsersQ(params.cfg.DB().Clone()),
		Cancel:          params.cancel,
		Clients:         params.clients,
		CertIntegrators: params.certIntegrators,
		dailyStorage:    storage.DailyStorageInstance(params.ctx),
		wg:              params.wg,
	}
}
