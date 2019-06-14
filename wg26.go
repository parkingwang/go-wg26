package wg26

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

//
// Author: 陈哈哈 bitschen@163.com
//

const (
	LengthNumber = 10
	LengthWG26SN = 8
)

var (
	bOrder = binary.BigEndian
)

// 维根26编码
type Wg26Id struct {
	CardSN       string // 原始十位卡号
	Wg26Hex      string // 十六进制
	Wg26Start    uint16 // 前段
	Wg26End      uint16 // 后段
	Wg26SN       string // WG26标准卡号
	Wg26SNFormat string // WG26标准8位卡号
}

// 从标准10位卡号解析
func ParseFromCardNumber(number string) *Wg26Id {
	nInt := TrimZero(number)
	nHex := fmt.Sprintf("%06X", nInt)
	bytes, _ := hex.DecodeString(nHex)
	start := uint16(bytes[0])
	end := bOrder.Uint16(bytes[1:])
	return &Wg26Id{
		CardSN:       number,
		Wg26Hex:      nHex,
		Wg26Start:    start,
		Wg26End:      end,
		Wg26SN:       fmt.Sprintf("%d%d", start, end),
		Wg26SNFormat: AppendZero(fmt.Sprintf("%d,%d", start, end), LengthWG26SN+1),
	}
}

// 从维根26国际标准编码解析
// [0] Start
// [1-2] End
func ParseFromWg26(wg26std [3]byte) *Wg26Id {
	start := uint16(wg26std[0])
	end := bOrder.Uint16(wg26std[1:])
	return &Wg26Id{
		CardSN:       AppendZero(fmt.Sprintf("%d", bOrder.Uint32([]byte{0, wg26std[0], wg26std[1], wg26std[2]})), LengthNumber),
		Wg26Hex:      fmt.Sprintf("%06X", wg26std),
		Wg26Start:    start,
		Wg26End:      end,
		Wg26SN:       fmt.Sprintf("%d%d", start, end),
		Wg26SNFormat: AppendZero(fmt.Sprintf("%d,%d", start, end), LengthWG26SN+1),
	}
}

// 从维根26标准字面卡号解析
func ParseFromWg26Number(wg26Number string) *Wg26Id {
	std := fmt.Sprintf("%8s", wg26Number)
	bytes := make([]byte, 2)
	bOrder.PutUint16(bytes, toInt(std[3:]))
	return ParseFromWg26([3]byte{
		byte(toInt(strings.TrimLeft(std[:3], " "))),
		bytes[0],
		bytes[1]})
}

// 从标准10位卡号的整数卡号解析
func ParseFromCardNumberValue(number uint32) *Wg26Id {
	b := make([]byte, 4)
	bOrder.PutUint32(b, number)
	return ParseFromWg26([3]byte{b[1], b[2], b[3]})
}

// TrimZero 删除前导0字符
func TrimZero(card string) int64 {
	card = strings.TrimLeft(card, "0")
	v, _ := strconv.ParseInt(card, 10, 64)
	return v
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

func toInt(v string) uint16 {
	r, _ := strconv.Atoi(v)
	return uint16(r)
}
