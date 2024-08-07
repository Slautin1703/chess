package usecases

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func AddPlayerToQueue(rdb *redis.Client, playerID string) error {
	return rdb.RPush(ctx, "player_queue", playerID).Err()
}

func GetNextPlayerFromQueue(rdb *redis.Client) (string, error) {
	playerID, err := rdb.LPop(ctx, "player_queue").Result()
	if err != nil {
		return "", err
	}
	return playerID, nil
}
