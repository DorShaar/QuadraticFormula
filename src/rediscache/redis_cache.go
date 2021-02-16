package rediscache

import (
	"log"
	"strconv"
	"github.com/gomodule/redigo/redis"
)

// https://www.alexedwards.net/blog/working-with-redis
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

func GetKey(a int, b int, c int) string {
	return strconv.Itoa(a) + "-" + strconv.Itoa(b) + "-" + strconv.Itoa(c) 
}

func (redisCache *RedisCache) GetRootsIfExist(key string) (float64, float64, bool) {
	if !redisCache.isConnected {
		log.Panicf("Could not get roots since redis cache is not connected")
	}

	log.Printf("Searching if there is an existing result for %s", key)

	root1, err := redis.Float64(redisCache.connection.Do("HGET", key, "root1"))

	if err != nil {
		log.Printf("Could not get key: %s root1", key)
		return 0, 0, false
	}

	root2, err := redis.Float64(redisCache.connection.Do("HGET", key, "root2"))

	if err != nil {
		log.Printf("Could not get key: %s root2", key)
		return 0, 0, false
	}

	return root1, root2, true
}

func (redisCache *RedisCache) SaveResult(key string, root1 float64, root2 float64) error {
	if !redisCache.isConnected {
		log.Panicf("Could not save result since redis cache is not connected")
	}

	_, err := redisCache.connection.Do("HMSET", key, "root1", root1, "root2", root2)

	if err != nil {
		log.Printf("Could not set key: %s", key)
		return err
	}

	return nil
}