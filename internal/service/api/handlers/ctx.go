package handlers

import (
	"context"
	"net/http"

	"github.com/dov-id/cert-integrator-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	contractsCtxKey
	DbCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func ContractsQ(r *http.Request) data.Contracts {
	return r.Context().Value(contractsCtxKey).(data.Contracts).New()
}

func CtxContractsQ(entry data.Contracts) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, contractsCtxKey, entry)
	}
}

func DB(r *http.Request) *pgdb.DB {
	return r.Context().Value(DbCtxKey).(*pgdb.DB)
}

func CtxDB(entry *pgdb.DB) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, DbCtxKey, entry)
	}
}
