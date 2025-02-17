package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"

	"github.com/Sutdown/go_study/mod/models"
)

// 具体操作

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
}

// InsertUser 想数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 生成加密密码
	password := encryptPassword([]byte(user.Password))
	// 把用户插入数据库
	sqlStr := "insert into user(user_id, username, password) values (?,?,?)"
	_, err = db.Exec(sqlStr, user.UserID, user.Username, password)
	return
}

const secret = "xiangfei01"

func encryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(secret))
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
