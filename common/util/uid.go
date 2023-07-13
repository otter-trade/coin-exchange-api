package util

import (
	"bytes"
	"strings"
)

/*
用于UID与6位邀请码的相互转化
34^6 = 1544804416 , 应该够用, 如果不够用了再增加邀请码长度
*/
const (
	InviteCodeMod         = 1544804377                           // 质数，比34^6略小
	InviteCodeReplaceKey  = 1298498081                           // 随机参数, 小于InviteCodePrime
	InviteCodeCharSet     = "7JC2DN4HA5MGI6BQZV3Y9W8UEORLKSTFPX" // 不含0，1, 和O,l 容易混淆
	InviteCodeCharSetSize = int64(len(InviteCodeCharSet))
)

var charToInt map[rune]int
var inviteCodeReplaceKeyInv uint64

func init() {
	charToInt = make(map[rune]int)
	for i, v := range InviteCodeCharSet {
		charToInt[v] = i
	}
	inviteCodeReplaceKeyInv = PowMod(InviteCodeReplaceKey, InviteCodeMod-2, InviteCodeMod)
}

func PowMod(x, n, mod uint64) uint64 {
	x %= mod
	ret := uint64(1)
	tmp := x
	for n > 0 {
		if n&1 == 1 {
			ret = ret * tmp % mod
		}
		tmp = tmp * tmp % mod
		n >>= 1
	}
	return ret
}

func UidToInviteCode(uid int64) string {
	if uid <= 0 || uid >= InviteCodeMod {
		return ""
	}
	value := uid * InviteCodeReplaceKey % InviteCodeMod
	ret := bytes.NewBuffer(nil)
	for i := 0; i < 6; i++ {
		c := value % InviteCodeCharSetSize
		value /= InviteCodeCharSetSize
		ret.WriteRune(rune(InviteCodeCharSet[c]))
	}
	return ret.String()
}

func InviteCodeToUid(code string) int64 {
	// 增强容错性,防止眼花,生成的邀请码中不含0,1
	code = strings.Replace(code, "0", "O", -1)
	code = strings.Replace(code, "1", "I", -1)
	code = strings.ToUpper(code)
	if len(code) != 6 {
		return 0
	}
	ret := int64(0)
	for i := 5; i >= 0; i-- {
		if v, ok := charToInt[rune(code[i])]; !ok {
			return 0
		} else {
			ret = ret*InviteCodeCharSetSize + int64(v)
		}
	}
	return ret * int64(inviteCodeReplaceKeyInv) % InviteCodeMod
}
