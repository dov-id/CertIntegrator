package handlers

import (
	"net/http"
	"strings"

	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/data/postgres"
	"github.com/dov-id/cert-integrator-svc/internal/service/api/requests"
	"github.com/dov-id/cert-integrator-svc/internal/service/api/responses"
	"github.com/ethereum/go-ethereum/common"
	"github.com/iden3/go-merkletree-sql/v2"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetPublicKeys(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetPublicKeysRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if request.Course == nil {
		Log(r).WithError(err).Error("course address is empty")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	course := *request.Course

	if request.Signer == nil {
		Log(r).WithError(err).Error("signer address is empty")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	signer := *request.Signer

	contract, err := ContractsQ(r).FilterByAddresses(course).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get course by address")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if contract == nil {
		Log(r).WithError(err).Errorf("no course with address `%s`", course)
		w.WriteHeader(http.StatusNotFound)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	treeStorage := postgres.NewStorage(Cfg(r).DB().Clone(), contract.Id)
	mTree, err := merkletree.NewMerkleTree(r.Context(), treeStorage, data.MaxMTreeLevel)
	if err != nil {
		Log(r).WithError(err).Error("failed to open/create merkle tree")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	users, err := UsersQ(r).Select()
	if err != nil {
		Log(r).WithError(err).Error("failed to get user")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	var participants = make([]data.User, 0)

	//TODO: make table `participants`, that will be store `user - course`
	for i := 0; i < len(users); i++ {
		if strings.ToLower(users[i].Address) == strings.ToLower(signer) {
			continue
		}

		keyBig := common.HexToAddress(users[i].Address).Big()
		//just check if user in tree (if user is participant)
		_, _, _, err = mTree.Get(r.Context(), keyBig)
		if err == merkletree.ErrKeyNotFound {
			continue
		}

		if err != nil && err != merkletree.ErrKeyNotFound {
			Log(r).WithError(err).Error("failed to get leaf")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		participants = append(participants, users[i])
	}

	w.WriteHeader(http.StatusOK)
	ape.Render(w, responses.NewUserListResponse(participants))
	return
}
