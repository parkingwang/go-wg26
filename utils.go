package wg26

import (
	"strconv"
	"strings"
)

//
// Author: 陈哈哈 bitschen@163.com
//

// TrimZeroToInt64 删除前导0字符
func TrimZeroToInt64(card string) int64 {
	card = strings.TrimLeft(card, "0")
	return ToInt64(card)
}

// AppendZero 添加前导0字符
func AppendZero(txt string, max int) string {
	s := max - len(txt)
	zeros := ""
	for i := 0; i < s; i++ {
		zeros += "0"
	}
	return zeros + txt
}

// IsDigits 返回字符串是否全为数字
func IsDigits(str string) bool {
	return strings.IndexFunc(str, func(c rune) bool {
		return c < '0' || c > '9'
	}) == -1
}

// IsCardSN 返回字符串是否为卡号。卡号包括前导0字符。
func IsCardSN(card string) bool {
	return LengthCardSN == len(card) && IsDigits(card)
}

func ToInt64(v string) int64 {
	r, _ := strconv.ParseInt(v, 10, 64)
	return r
}
