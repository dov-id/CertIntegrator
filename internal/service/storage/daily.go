package storage

import (
	"context"
	"sync"
	"time"
)

const dailyStorageConst = "DailyStorage"

type DailyStorage interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
	Delete(key string)
	Clear()

	Run(ctx context.Context)
}

type dailyStorage struct {
	mapping sync.Map
}

func NewDailyStorage() DailyStorage {
	return &dailyStorage{}
}

func (ds *dailyStorage) Get(key string) (interface{}, bool) {
	val, ok := ds.mapping.Load(key)
	return val, ok
}

func (ds *dailyStorage) Set(key string, value interface{}) {
	ds.mapping.Store(key, value)
}

func (ds *dailyStorage) Delete(key string) {
	ds.mapping.Delete(key)
}

func (ds *dailyStorage) Clear() {
	ds.mapping.Range(func(key interface{}, value interface{}) bool {
		ds.mapping.Delete(key)
		return true
	})
}

func (ds *dailyStorage) Run(ctx context.Context) {
	for {
		timer := initTimer()
		select {
		case <-timer.C:
			ds.Clear()
		case <-ctx.Done():
			timer.Stop()
			break
		}
	}
}

// must process at 00:00 AM every day
func initTimer() *time.Timer {
	now := time.Now().UTC()
	nextDay := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC)

	return time.NewTimer(time.Until(nextDay))
}

func DailyStorageInstance(ctx context.Context) DailyStorage {
	return ctx.Value(dailyStorageConst).(DailyStorage)
}

func CtxDailyStorage(entry DailyStorage, ctx context.Context) context.Context {
	return context.WithValue(ctx, dailyStorageConst, entry)
}
