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
		Key:   utils.RandomKey(),
		Value: utils.RandomValue(),
	}
	cache, err := testQueries.Set(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, cache.Key, arg.Key)
	require.Equal(t, cache.Value, arg.Value)
	return cache
}

func TestSet(t *testing.T) {
	cacheRandomEntry(t)
}

func TestGet(t *testing.T) {
	entry := cacheRandomEntry(t)

	cache, err := testQueries.Get(context.Background(), entry.Key)
	require.NoError(t, err)
	require.Equal(t, entry.Key, cache.Key)
	require.Equal(t, entry.Value, cache.Value)
	require.Equal(t, entry.Ttl, cache.Ttl)
}

func TestSetUpdate(t *testing.T) {
	arg := SetParams{
		Key:   utils.RandomKey(),
		Value: utils.RandomValue(),
	}
	cache, err := testQueries.Set(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, cache.Key, arg.Key)
	require.Equal(t, cache.Value, arg.Value)

	// update
	arg.Value = utils.RandomValue()
	updatedCache, err := testQueries.Set(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, updatedCache.Key, cache.Key)
	require.Equal(t, updatedCache.Value, arg.Value)
	require.NotEqual(t, updatedCache.Value, cache.Value)
}

func TestDelete(t *testing.T) {
	cache := cacheRandomEntry(t)
	err := testQueries.Delete(context.Background(), cache.Key)
	require.NoError(t, err)
}

func BenchmarkSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arg := SetParams{
			Key:   utils.RandomKey(),
			Value: utils.RandomValue(),
			Ttl: sql.NullInt32{
				Int32: utils.RandomTTl(),
				Valid: true,
			},
		}
		_, _ = testQueries.Set(context.Background(), arg)
	}
}
