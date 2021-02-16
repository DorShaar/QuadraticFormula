package rediscachetests

import (
	"rediscache"
	"testing"
)

func TestFindRoots_EquationWithRoots_RootsAreCorrect(t *testing.T) {
	redisConnectionAddress := "localhost:6379"
	redisCache := rediscache.RedisCache{}

	redisCache.Connect(redisConnectionAddress)

	key := rediscache.GetKey(3, 2, 1)

	_, _, exists := redisCache.GetRootsIfExist(key)

	if exists {
		t.FailNow()
	}

	err := redisCache.SaveResult(key, 4.5, -2)

	if err != nil {
		t.FailNow()
	}

	root1, root2, exists := redisCache.GetRootsIfExist(key)

	if !exists {
		t.FailNow()
	}

	if root1 != 4.5 && root2 != -2 {
		t.FailNow()
	}
}