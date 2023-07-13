package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassWord(password string) (hashedPassword string, err error) {
	PassWord := []byte(password)
	Password, err := bcrypt.GenerateFromPassword(PassWord, bcrypt.DefaultCost)
	hashedPassword = string(Password)
	return
}

func CheckPassWord(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

/*
*
给密码hash加密
*/
func PasswordHash(password, salt string) string {
	hashPwd, _ := bcrypt.GenerateFromPassword([]byte(salt+password), bcrypt.MinCost)
	return string(hashPwd)
}

/*
*
检测用户上传的密码是否正确
*/
func PasswordCheck(hashPwd, salt, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(salt+password))
	if err != nil {
		return false
	}
	return true
}

// HashMake
func HashMake(plainPwd string) (hashedPwd string, err error) {
	var hashed []byte
	hashed, err = bcrypt.GenerateFromPassword([]byte(plainPwd), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	hashedPwd = string(hashed)
	return
}

// HashIsSame
func HashIsSame(plainPwd, hashedPwd string) (yes bool, err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		return
	}
	yes = true
	return
}

// HashNeedRefresh
func HashNeedRefresh(hashedPwd string) bool {
	hashCost, err := bcrypt.Cost([]byte(hashedPwd))
	return err != nil || hashCost != bcrypt.DefaultCost
}
