package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Config interface {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer

	Timeouts() *TimeoutsCfg
	PublicKeyRetriever() *PublicKeyRetrieverCfg
	Networks() *NetworksCfg
	Wallet() *WalletCfg
	RpcProvider() *RpcProviderCfg
	CertificatesIssuer() *ContractsCfg
	CertificatesFabric() *ContractsCfg
	CertificatesIntegrator() *CertificatesIntegratorCfg
}

type config struct {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer
	getter kv.Getter

	publicKeyRetriever     comfig.Once
	networks               comfig.Once
	wallet                 comfig.Once
	timeouts               comfig.Once
	rpcProvider            comfig.Once
	certificatesIssuer     comfig.Once
	certificatesFabric     comfig.Once
	certificatesIntegrator comfig.Once
}

func New(getter kv.Getter) Config {
	return &config{
		getter:     getter,
		Databaser:  pgdb.NewDatabaser(getter),
		Copuser:    copus.NewCopuser(getter),
		Listenerer: comfig.NewListenerer(getter),
		Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),
	}
}
