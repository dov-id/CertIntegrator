package handlers

import (
	"net/http"

	"github.com/dov-id/cert-integrator-svc/internal/service/api/requests"
	"github.com/dov-id/cert-integrator-svc/internal/service/api/responses"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetUsersCount(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetUsersCountRequest(r)
	if err != nil {
		Log(r).WithError(err).Debug("bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	contract, err := MasterQ(r).ContractsQ().FilterByAddresses(*request.Course).Get()
	if err != nil {
		Log(r).WithError(err).Debug("failed to get course by address")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if contract == nil {
		Log(r).WithError(err).Debugf("no course with address `%s`", *request.Course)
		w.WriteHeader(http.StatusNotFound)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	users, err := MasterQ(r).UsersQ().FilterByContractId(contract.Id).Select()
	if err != nil {
		Log(r).WithError(err).Debug("failed to select users")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
	ape.Render(w, responses.NewCountResponse(int64(len(users))))
	return
}
