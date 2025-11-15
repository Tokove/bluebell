package redis

import "errors"

var (
	ErrorVoteTimeExpire = errors.New("投票时间已过")
)
