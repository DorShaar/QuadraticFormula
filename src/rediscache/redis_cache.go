package rediscache

import (
	"log"
	"github.com/gomodule/redigo/redis"
)

type RedisCache struct {
	connection 		redis.Conn
	isConnected 	bool
}

// Connect Creates a connection to given address
func (redisCache *RedisCache) Connect(redisConnectionAddress string) {
	connection, err := redis.Dial("tcp", redisConnectionAddress)

	if err != nil {
		log.Fatalf("Could not connect to %s", redisConnectionAddress)
		return
	}

	redisCache.isConnected = true
	redisCache.connection = connection

	log.Printf("Connected to address: %s", redisConnectionAddress)
}

// Disconnect closes that connection
func (redisCache *RedisCache) Disconnect() {
	redisCache.isConnected = false
	redisCache.connection.Close()

	log.Printf("Disconnected")
}

func (redisCache *RedisCache) GetRootsIfExist(key string) (int, int, int, bool) {
	if !redisCache.isConnected {
		log.Panicf("Could not get roots since redis cache is not connected")
	}

	log.Printf("Searching if there is an existing result for %s", key)

	a, err := redis.Int(redisCache.connection.Do("HGET", key, "a"))

	if err != nil {
		log.Printf("Could not get key: %s", key)
		return 0, 0, 0, false
	}

	b, err := redis.Int(redisCache.connection.Do("HGET", key, "b"))

	if err != nil {
		log.Printf("Could not get key: %s", key)
		return 0, 0, 0, false
	}

	c, err := redis.Int(redisCache.connection.Do("HGET", key, "c"))

	if err != nil {
		log.Printf("Could not get key: %s", key)
		return 0, 0, 0, false
	}

	return a, b, c, true
}

func (redisCache *RedisCache) SaveResult(key string, a int, b int, c int) error {
	if !redisCache.isConnected {
		log.Panicf("Could not get roots since redis cache is not connected")
	}

	_, err := redisCache.connection.Do("HSET", key, "a", a, "b", b, "c", c)

	if err != nil {
		log.Printf("Could not set key: %s", key)
		return err
	}

	return nil
}