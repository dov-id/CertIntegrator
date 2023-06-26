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
	mapping map[string]interface{}
	mutex   sync.Mutex
}

func NewDailyStorage() DailyStorage {
	return &dailyStorage{
		mapping: make(map[string]interface{}),
	}
}

func (ds *dailyStorage) Get(key string) (interface{}, bool) {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()
	val, ok := ds.mapping[key]
	return val, ok
}

func (ds *dailyStorage) Set(key string, value interface{}) {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()
	ds.mapping[key] = value
}

func (ds *dailyStorage) Delete(key string) {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()
	delete(ds.mapping, key)
}

func (ds *dailyStorage) Clear() {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()
	ds.mapping = make(map[string]interface{})
}

func (ds *dailyStorage) Run(ctx context.Context) {
	go func() {
		for {
			now := time.Now().UTC()
			// must process at 00:00 AM every day
			nextDay := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC)
			timer := time.NewTimer(time.Until(nextDay))
			select {
			case <-timer.C:
				ds.Clear()
			case <-ctx.Done():
				timer.Stop()
				break
			}
		}
	}()
}

func DailyStorageInstance(ctx context.Context) DailyStorage {
	return ctx.Value(dailyStorageConst).(DailyStorage)
}

func CtxDailyStorage(entry DailyStorage, ctx context.Context) context.Context {
	return context.WithValue(ctx, dailyStorageConst, entry)
}
