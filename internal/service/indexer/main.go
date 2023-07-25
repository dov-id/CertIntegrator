package indexer

import (
	"context"

	"github.com/dov-id/cert-integrator-svc/contracts"
	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/data/postgres"
	"github.com/dov-id/cert-integrator-svc/internal/service/sender"
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
	Type string

	cfg config.Config
	log *logan.Entry

	issuerCh  chan string
	Addresses []string
	Blocks    map[string]int64

	ContractsQ    data.Contracts
	UsersQ        data.Users
	ParticipantsQ data.Participants
	TransactionsQ data.Transactions

	Clients         map[types.Network]*ethclient.Client
	CertIntegrators map[types.Network]*contracts.CertIntegratorContract

	dailyStorage storage.DailyStorage
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
	clients         map[types.Network]*ethclient.Client
	certIntegrators map[types.Network]*contracts.CertIntegratorContract
}

func Run(cfg config.Config, ctx context.Context) {
	params, err := prepareIndexerParams(cfg, ctx)
	if err != nil {
		panic(errors.Wrap(err, "failed to prepare indexer params"))
	}

	params.name = IssuerContract
	NewIndexer(*params).Run(ctx)

	params.name = FabricContract
	NewIndexer(*params).Run(ctx)

	//TODO: is it okay to run runner from this runner? It's for not initializing contracts and clients twice
	sender.NewSender(cfg, params.clients, params.certIntegrators).Run(ctx)
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
		ContractsQ:      postgres.NewContractsQ(params.cfg.DB().Clone()),
		UsersQ:          postgres.NewUsersQ(params.cfg.DB().Clone()),
		ParticipantsQ:   postgres.NewParticipantsQ(params.cfg.DB().Clone()),
		TransactionsQ:   postgres.NewTransactionsQ(params.cfg.DB().Clone()),
		Clients:         params.clients,
		CertIntegrators: params.certIntegrators,
		dailyStorage:    storage.DailyStorageInstance(params.ctx),
	}
}
