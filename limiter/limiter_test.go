package limiter

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ulule/limiter/v3"
	limiterredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

func TestLimiter(t *testing.T) {
	redisServer := miniredis.NewMiniRedis()
	require.NoError(t, redisServer.Start())
	redisCli := redis.NewClient(&redis.Options{
		Addr: redisServer.Addr(),
	})
	t.Cleanup(func() {
		require.NoError(t, redisCli.Close())
	})

	rate := limiter.Rate{
		Limit:  5,
		Period: 5 * time.Second,
	}
	store, err := limiterredis.NewStoreWithOptions(redisCli, limiter.StoreOptions{
		Prefix: "test",
	})

	assert.NoError(t, err)

	for i := 0; i < 10; i++ {
		expect := false
		if i >= 5 {
			expect = true
		}

		res1, err := store.Get(context.Background(), "key1", rate)
		assert.NoError(t, err)
		assert.Equal(t, expect, res1.Reached)

		res2, err := store.Get(context.Background(), "key2", rate)
		assert.NoError(t, err)
		assert.Equal(t, expect, res2.Reached)
	}
	redisServer.FastForward(6 * time.Second)
	res1, err := store.Get(context.Background(), "key1", rate)
	assert.NoError(t, err)
	assert.Equal(t, false, res1.Reached)
}
