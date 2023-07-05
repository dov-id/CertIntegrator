package handlers

import (
	"context"
	"net/http"

	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	contractsCtxKey
	usersCtxKey
	CfgCtxKey
	ParentCtxCtxKey
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

func UsersQ(r *http.Request) data.Users {
	return r.Context().Value(usersCtxKey).(data.Users).New()
}

func CtxUsersQ(entry data.Users) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, usersCtxKey, entry)
	}
}

func Cfg(r *http.Request) config.Config {
	return r.Context().Value(CfgCtxKey).(config.Config)
}

func CtxCfg(entry config.Config) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, CfgCtxKey, entry)
	}
}

func ParentCtx(r *http.Request) context.Context {
	return r.Context().Value(ParentCtxCtxKey).(context.Context)
}

func CtxParentCtx(entry context.Context) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ParentCtxCtxKey, entry)
	}
}
