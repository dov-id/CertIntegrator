package handlers

import (
	"net/http"

	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/service/api/responses"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetCourses(w http.ResponseWriter, r *http.Request) {
	contracts, err := ContractsQ(r).FilterByTypes(data.ISSUER).Select()
	if err != nil {
		Log(r).WithError(err).Errorf("failed to select courses")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
	ape.Render(w, responses.NewCourseListResponse(contracts))
}
