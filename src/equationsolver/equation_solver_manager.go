package equationsolver

import (
	"log"
	"rediscache"
)

type EquationSolverManager struct {
	cache		rediscache.RedisCache
}

func (equationSolverManager *EquationSolverManager) Init(redisConnectionAddress string) {
	equationSolverManager.cache.Connect(redisConnectionAddress)
}

func (equationSolverManager *EquationSolverManager) FindRoots(a int, b int, c int) (float64, float64, bool) {
	key := rediscache.GetKey(a, b, c)
	root1, root2, exists := equationSolverManager.cache.GetRootsIfExist(key)

	if exists {
		log.Printf("Found roots from cache. Key: %s, root1: %f, root2: %f", key, root1, root2)
		return root1, root2, true
	}

	root1, root2, hasRoots := CalculateRoots(a, b, c)
	if !hasRoots {
		log.Printf("There are no roots for a %d b %d c %d", a, b, c)
		return 0, 0, false
	}

	log.Printf("The roots a %d b %d c %d are: {%f, %f}", a, b, c, root1, root2)
	return root1, root2, true
}