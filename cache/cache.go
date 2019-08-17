package cache

import (
	"time"

	aw "github.com/deanishe/awgo"
)

const (
	maxCacheAgeMinutes = 1 // How long to cache data
)

type BaseCache struct {
	wf *aw.Workflow
}

func getMaxCacheAge() time.Duration {
	return maxCacheAgeMinutes * time.Minute
}

func (cache *BaseCache) getRawCache(cacheKey string, expire time.Duration, v interface{}) error {
	store := cache.wf.Cache

	if !store.Exists(cacheKey) {
		return nil
	}

	if store.Expired(cacheKey, expire) {
		return nil
	}

	if err := store.LoadJSON(cacheKey, &v); err != nil {
		return err
	}

	return nil
}

func (cache *BaseCache) storeRawData(cacheKey string, v interface{}) (interface{}, error) {
	store := cache.wf.Cache

	if err := store.StoreJSON(cacheKey, v); err != nil {
		return nil, err
	}

	return v, nil
}
