package redis

// redis keys

const (
	KeyPrefix              = "bluebell:"
	KeyPostTimeZSet        = "post:time"  // zset; 帖子及发帖时间
	KeyPostScoreZSet       = "post:score" // zset; 帖子及帖子分数
	KeyPostVotedZSetPrefix = "post:voted" // zset; 记录用户及投票类型; 参数是post_id 后面要接postID
)

// 给 key 加前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
