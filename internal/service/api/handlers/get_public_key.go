package handlers

import (
	"net/http"

	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/helpers"
	"github.com/dov-id/cert-integrator-svc/internal/service/api/requests"
	"github.com/dov-id/cert-integrator-svc/internal/service/api/responses"
	"github.com/ethereum/go-ethereum/common"
	pkgErrors "github.com/pkg/errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func GetPublicKey(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetPublicKeyRequest(r)
	if err != nil {
		Log(r).WithError(err).Debug("bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	contract, err := MasterQ(r).ContractsQ().FilterByAddresses(*request.Course).Get()
	if err != nil {
		Log(r).WithError(err).Debug("failed to get contract")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if contract == nil {
		Log(r).WithError(err).Debug("failed to get contract")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	user, err := MasterQ(r).UsersQ().FilterByAddresses(*request.Address).Get()
	if err != nil {
		Log(r).WithError(err).Debug("failed to get user")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if user != nil {
		w.WriteHeader(http.StatusOK)
		ape.Render(w, responses.NewUserResponse(*user))
		return
	}

	err = findPubKey(r, *request.Address, int64(contract.Id))
	if err != nil {
		if pkgErrors.Is(err, data.ErrNoPublicKey) {
			w.WriteHeader(http.StatusNotFound)
			ape.Render(w, problems.NotFound())
			return
		}

		Log(r).WithError(err).Debug("failed to find public key")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	user, err = MasterQ(r).UsersQ().FilterByAddresses(*request.Address).Get()
	if err != nil {
		Log(r).WithError(err).Debug("failed to get user")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if user != nil {
		w.WriteHeader(http.StatusOK)
		ape.Render(w, responses.NewUserResponse(*user))
		return
	}

	w.WriteHeader(http.StatusNotFound)
	ape.Render(w, problems.NotFound())
	return
}

func findPubKey(r *http.Request, address string, contractId int64) error {
	clients, err := helpers.GetNetworkClients(ParentCtx(r))
	if err != nil {
		return errors.Wrap(err, "failed to init network clents")
	}

	err = helpers.ProcessPublicKey(helpers.ProcessPubKeyParams{
		Ctx:        r.Context(),
		Cfg:        Cfg(r),
		Address:    common.HexToAddress(address),
		UsersQ:     MasterQ(r).UsersQ(),
		ContractId: uint64(contractId),
		Clients:    clients,
	})
	if err != nil {
		return errors.Wrap(err, "failed to process public key")
	}

	return nil
}
