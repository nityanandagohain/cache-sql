package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/nityanandagohain/sql-cache/db/utils"
	"github.com/stretchr/testify/require"
)

func cacheRandomEntry(t *testing.T) Cache {
	arg := SetParams{
		CacheKey: utils.RandomKey(),
		Value:    utils.RandomValue(),
	}
	err := testQueries.Set(context.Background(), arg)
	require.NoError(t, err)
	cache, err := testQueries.Get(context.Background(), arg.CacheKey)
	require.NoError(t, err)
	require.Equal(t, cache.CacheKey, arg.CacheKey)
	require.Equal(t, cache.Value, arg.Value)
	return cache
}

func TestSet(t *testing.T) {
	cacheRandomEntry(t)
}

func TestGet(t *testing.T) {
	entry := cacheRandomEntry(t)

	cache, err := testQueries.Get(context.Background(), entry.CacheKey)
	require.NoError(t, err)
	require.Equal(t, entry.CacheKey, cache.CacheKey)
	require.Equal(t, entry.Value, cache.Value)
	require.Equal(t, entry.Ttl, cache.Ttl)
}

func TestCaching(t *testing.T) {
	entry := cacheRandomEntry(t)

	cache, err := testQueries.Get(context.Background(), entry.CacheKey)
	require.NoError(t, err)
	require.Equal(t, entry.CacheKey, cache.CacheKey)
	require.Equal(t, entry.Value, cache.Value)
	require.Equal(t, entry.Ttl, cache.Ttl)
}

func TestSetUpdate(t *testing.T) {
	arg := SetParams{
		CacheKey: utils.RandomKey(),
		Value:    utils.RandomValue(),
	}
	err := testQueries.Set(context.Background(), arg)
	require.NoError(t, err)
	cache, err := testQueries.Get(context.Background(), arg.CacheKey)
	require.NoError(t, err)
	require.Equal(t, cache.CacheKey, arg.CacheKey)
	require.Equal(t, cache.Value, arg.Value)

	// update
	// arg.Value = utils.RandomValue()
	// err = testQueries.Set(context.Background(), arg)
	// require.NoError(t, err)
	// require.NoError(t, err)
	// updatedCache, err := testQueries.Get(context.Background(), arg.CacheKey)
	// require.NoError(t, err)
	// require.Equal(t, updatedCache.CacheKey, cache.CacheKey)
	// require.Equal(t, updatedCache.Value, arg.Value)
	// require.NotEqual(t, updatedCache.Ttl, cache.Ttl)
}

func TestDelete(t *testing.T) {
	cache := cacheRandomEntry(t)
	err := testQueries.Delete(context.Background(), cache.CacheKey)
	require.NoError(t, err)
}

func BenchmarkSet(b *testing.B) {
	arg := SetParams{
		CacheKey: utils.RandomKey(),
		Value:    utils.RandomValue(),
		Ttl: sql.NullInt32{
			Int32: utils.RandomTTl(),
			Valid: true,
		},
	}
	err := testQueries.Set(context.Background(), arg)
	require.NoError(b, err)
	for i := 0; i < b.N; i++ {
		_, err := testQueries.Get(context.Background(), arg.CacheKey)
		require.NoError(b, err)
	}
}
