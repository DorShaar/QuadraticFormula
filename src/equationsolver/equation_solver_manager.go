package equationsolver

import (
	"log"
	"rediscache"
)

type EquationSolverManager struct {
	cache		rediscache.RedisCache
}

func (equationSolverManager *EquationSolverManager) Init() {
	// TODO
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
		return 0, 0, false
	}

	return root1, root2, true
}

        // public Task SendMessage(EquationRoots equationRootsMessage)
        // {
        //     string serializedString = JsonSerializer.Serialize(equationRootsMessage);
        //     mLogger.LogInformation($"Sending message: {serializedString}");

        //     return Task.CompletedTask;
        // }