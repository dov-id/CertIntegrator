package handlers

import (
	"net/http"

	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/service/api/responses"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetCourses(w http.ResponseWriter, r *http.Request) {
	contracts, err := MasterQ(r).ContractsQ().FilterByTypes(data.Issuer).Select()
	if err != nil {
		Log(r).WithError(err).Debugf("failed to select courses")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
	ape.Render(w, responses.NewCourseListResponse(contracts))
}
