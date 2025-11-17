package redis

import (
	"bluebell_backend/model"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func CreatePost(postID, communityID uint64) error {
	now := float64(time.Now().Unix())
	pipeline := client.TxPipeline()
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  now,
		Member: postID,
	})
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  now,
		Member: postID,
	})
	cKeyTime := getRedisKey(KeyPostTimeZSet + strconv.FormatUint(communityID, 10))
	cKeyScore := getRedisKey(KeyPostScoreZSet + strconv.FormatUint(communityID, 10))
	pipeline.ZAdd(cKeyTime, redis.Z{
		Score:  now,
		Member: postID,
	})
	pipeline.ZAdd(cKeyScore, redis.Z{
		Score:  now,
		Member: postID,
	})
	_, err := pipeline.Exec()
	return err
}

func GetPostIDsInOrder(p *model.ParamPostList) ([]string, error) {
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == model.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	return client.ZRevRange(key, (p.Page-1)*p.Size, p.Page*p.Size-1).Result()
}

func GetPostVoteData(ids []string) (data []int64, err error) {
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPrefix + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

func GetCommunityPostIDsInOrder(p *model.ParamPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == model.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}
	key := orderKey + strconv.FormatUint(p.CommunityID, 10)
	if client.Exists(key).Val() < 1 {
		cKey := getRedisKey(KeyCommunitySetPrefix + strconv.FormatUint(p.CommunityID, 10))
		pipeline := client.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, cKey, orderKey)
		pipeline.Expire(key, 60*time.Second)
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	return client.ZRevRange(key, (p.Page-1)*p.Size, p.Page*p.Size-1).Result()
}
