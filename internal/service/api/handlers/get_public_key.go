package handlers

import (
	"context"
	"net/http"

	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/helpers"
	"github.com/dov-id/cert-integrator-svc/internal/service/api/models"
	"github.com/dov-id/cert-integrator-svc/internal/service/api/requests"
	"github.com/dov-id/cert-integrator-svc/internal/service/storage"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func GetPublicKey(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetPublicKeyRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if request.Address == nil {
		Log(r).WithError(err).Error("address is empty")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	address := *request.Address

	user, err := UsersQ(r).FilterByAddresses(address).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get user")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if user != nil {
		w.WriteHeader(http.StatusOK)
		ape.Render(w, models.NewUserResponse(*user))
		return
	}

	clients, err := helpers.InitNetworkClients(Cfg(r).Networks().Networks)
	if err != nil {
		Log(r).WithError(err).Error("failed to init network clents")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	err = helpers.ProcessPublicKey(helpers.ProcessPubKeyParams{
		Cfg:     Cfg(r),
		Address: common.HexToAddress(address),
		UsersQ:  UsersQ(r),
		Storage: storage.DailyStorageInstance(ParentCtx(r)),
		Clients: clients,
	})
	if err != nil {
		if err.Error() == data.NoPublicKeyErr {
			amount, err := getLeftAttemptsAmount(ParentCtx(r), Cfg(r), common.HexToAddress(address))
			if err != nil {
				Log(r).WithError(err).Error("failed to get left attempts amount")
				ape.RenderErr(w, problems.InternalError())
				return
			}

			w.WriteHeader(http.StatusNotFound)
			ape.Render(w, models.NewAttemptsResponse(amount))
			return
		}
		Log(r).WithError(err).Error("failed to process public key")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	for _, client := range clients {
		client.Close()
	}

	user, err = UsersQ(r).FilterByAddresses(address).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get user")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if user != nil {
		w.WriteHeader(http.StatusOK)
		ape.Render(w, models.NewUserResponse(*user))
		return
	}

	amount, err := getLeftAttemptsAmount(ParentCtx(r), Cfg(r), common.HexToAddress(address))
	if err != nil {
		Log(r).WithError(err).Error("failed to get left attempts amount")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNotFound)
	ape.Render(w, models.NewAttemptsResponse(amount))
	return
}

func getLeftAttemptsAmount(ctx context.Context, cfg config.Config, address common.Address) (int64, error) {
	amount, err := helpers.Get(storage.DailyStorageInstance(ctx), address)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get amount from storage")
	}

	return cfg.Attempts().Daily - amount, nil
}
