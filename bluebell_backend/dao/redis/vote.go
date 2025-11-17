package redis

import (
	"time"

	"github.com/go-redis/redis"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432
)

// direction = 1
// 1. 没投过， 现投赞成 1倍   +432     0 1 1
// 2. 投过反对， 现投赞成 2倍 +432*2  -1 1 2
// 2. 投过反对， 现取消 1倍   +432    -1 0 1
// direction = 0
// 1. 投赞成， 现取消 1倍     -432    1 0 -1
// direction = -1
// 1. 没投过， 现投反对 1倍   -432    0 -1 -1
// 2. 投过赞成， 现投反对 2倍 -432*2  1 -1 -2
func VoteForPost(userID, postID string, nowDirection float64) error {
	// 1. 查时间
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrorVoteTimeExpire
	}
	// 2. 查记录
	key := getRedisKey(KeyPostVotedZSetPrefix + postID)
	prevDirection := client.ZScore(key, userID).Val()
	diff := nowDirection - prevDirection
	if diff == 0{
		return ErrorVoteRepeat
	}
	// 修改后必须记录，用pipeline
	pipeline := client.TxPipeline()
	// 修改分数
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), diff*scorePerVote, postID)
	// 记录投票
	if nowDirection == 0 {
		pipeline.ZRem(key, userID)
	} else {
		pipeline.ZAdd(key, redis.Z{
			Score:  nowDirection,
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}

// redis 数据类型： string, hash, list, set, zset
// ZSet: 增 ZAdd 删 ZRem 改 ZIncrBy 查 ZScore
