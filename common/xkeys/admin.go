package xkeys

import "fmt"

const (
	loginToken = "console:token:%d"
)

// 账户信息
func LoginTokenKey(userId int64) string {
	return fmt.Sprintf(loginToken, userId)
}
