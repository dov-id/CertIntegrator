package handlers

import (
	"net/http"

	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/service/api/requests"
	"github.com/dov-id/cert-integrator-svc/internal/service/api/responses"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetUsersRequest(r)
	if err != nil {
		Log(r).WithError(err).Debug("bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	contract, err := MasterQ(r).ContractsQ().FilterByAddresses(request.Data.Attributes.Course).Get()
	if err != nil {
		Log(r).WithError(err).Debug("failed to get course by address")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if contract == nil {
		Log(r).WithError(err).Debugf("no course with address `%s`", request.Data.Attributes.Course)
		w.WriteHeader(http.StatusNotFound)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	users := make([]data.User, 0)

	for _, id := range request.Data.Attributes.Ids {
		user, err := MasterQ(r).UsersQ().FilterByContractId(contract.Id).Offset(uint64(id) - 1).Limit(1).Get()
		if err != nil {
			Log(r).WithError(err).Debug("failed to select users")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		if user == nil {
			Log(r).Debug("no such user")
			ape.RenderErr(w, problems.NotFound())
			return
		}

		users = append(users, *user)
	}

	w.WriteHeader(http.StatusOK)
	ape.Render(w, responses.NewUserListResponse(users))
	return
}
