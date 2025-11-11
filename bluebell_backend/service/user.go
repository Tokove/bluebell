package service

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/model"
	"bluebell_backend/pkg/crypto"
	"bluebell_backend/pkg/errcode"
	"bluebell_backend/pkg/jwt"
	"bluebell_backend/pkg/snowflake"
)

func Register(p *model.ParamSignUp) (err error) {
	// 查询用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 生成UID
	var userID uint64
	userID, err = snowflake.GetID()
	if err != nil {
		return err
	}
	// 加密
	user := &model.User{
		UserID:   userID,
		Username: p.Username,
	}
	user.Password, err = crypto.HashPassword(p.Password)
	if err != nil {
		return err
	}
	// 保存进数据库
	return mysql.InsertUser(user)
}

func Login(user *model.User) (string, error) {
	// 查询用户密码
	u, err := mysql.GetUserByUsername(user.Username)
	if err != nil {
		return "", err
	}
	// 比对密码
	if err := crypto.CheckPassword(u.Password, user.Password); err != nil {
		return "", errcode.ErrorInvalidPassword
	}
	return jwt.GenToken(u.UserID, u.Username)
}
