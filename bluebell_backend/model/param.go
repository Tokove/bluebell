package model

const (
	OrderTime   = "time"
	OrderScore  = "score"
	DefaultPage = 1
	DefaultSize = 10
)

type ParamSignUp struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVoteData struct {
	PostID    string `json:"post_id" binding:"required"`
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` // 1 0 -1 赞成 取消 反对
}

type ParamPostList struct {
	Page  int64  `form:"page"`
	Size  int64  `form:"size"`
	Order string `form:"order"`
	CommunityID uint64 `json:"community_id" form:"community_id"`
}
