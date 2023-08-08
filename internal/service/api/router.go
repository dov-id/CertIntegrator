package api

import (
	"github.com/dov-id/cert-integrator-svc/internal/data/postgres"
	"github.com/dov-id/cert-integrator-svc/internal/service/api/handlers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *Router) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxCfg(s.cfg),
			handlers.CtxParentCtx(s.ctx),
			handlers.CtxMasterQ(postgres.NewMasterQ(s.cfg.DB().Clone())),
		),
	)
	r.Route("/integrations/cert-integrator-svc", func(r chi.Router) {
		r.Post("/proof", handlers.GenerateSMTProof)

		r.Get("/courses", handlers.GetCourses)

		r.Route("/users", func(r chi.Router) {
			r.Post("/", handlers.GetUsers)
			r.Get("/count", handlers.GetUsersCount)
		})
	})

	return r
}
