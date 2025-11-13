package service

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/model"
	"bluebell_backend/pkg/crypto"
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

func Login(u *model.User) (*model.User, error) {
	// 查询用户密码
	user, err := mysql.GetUserByUsername(u.Username)
	if err != nil {
		return nil, mysql.ErrorUserNotExist
	}
	// 比对密码
	if err := crypto.CheckPassword(user.Password, u.Password); err != nil {
		return nil, mysql.ErrorInvalidPassword
	}
	// 生成token
	token, err := jwt.GenToken(u.UserID, u.Username)
	if err != nil {
		return nil, err
	}
	user.Token = token
	return user, nil
}
