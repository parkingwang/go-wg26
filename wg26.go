package wg26

import (
	"encoding/binary"
	"fmt"
	"strings"
)

//
// Author: 陈哈哈 bitschen@163.com
//

const (
	LengthCardSN = 10
	LengthWG26SN = 8
)

var (
	bigOrder = binary.BigEndian
)

// 维根26编码
type Wg26Id struct {
	CardSN    string // 原始十位卡号
	CardHex   string // 十六进制
	Wg26Start uint16 // 前段
	Wg26End   uint16 // 后段
	Wg26SN    string // WG26标准卡号
}

// ValueOfCardSN 返回卡号字面字符串的数值类型
func (wg *Wg26Id) ValueOfCardSN() uint32 {
	return uint32(ToInt64(wg.CardSN))
}

// ValueOfWg26SN 返回Wg26SN的字面字符串的数值类型
func (wg *Wg26Id) ValueOfWg26SN() uint32 {
	return uint32(ToInt64(wg.Wg26SN))
}

// FormatWg26SN 返回Wg26字面格式化的卡号字符串
func (wg *Wg26Id) FormatWg26SN() string {
	return AppendZero(fmt.Sprintf("%d,%d", wg.Wg26Start, wg.Wg26End), LengthWG26SN+1)
}

////

// 从标准10位卡号解析
func ParseFromCardNumber(number string) *Wg26Id {
	value := TrimZeroToInt64(number)
	return ParseFromCardNumberValue(uint32(value))
}

// ParseFromCardNumberValue 从标准10位卡号的整数卡号解析
func ParseFromCardNumberValue(number uint32) *Wg26Id {
	b := make([]byte, 4)
	bigOrder.PutUint32(b, number)
	return ParseFromWg26([3]byte{b[1], b[2], b[3]})
}

// ParseFromWg26 从维根26国际标准编码解析；其中字节为[0] Start，[1-2] End
func ParseFromWg26(wg26bytes [3]byte) *Wg26Id {
	start := uint16(wg26bytes[0])
	end := bigOrder.Uint16(wg26bytes[1:])
	return &Wg26Id{
		CardSN:    AppendZero(fmt.Sprintf("%d", bigOrder.Uint32([]byte{0, wg26bytes[0], wg26bytes[1], wg26bytes[2]})), LengthCardSN),
		CardHex:   fmt.Sprintf("%06X", wg26bytes),
		Wg26Start: start,
		Wg26End:   end,
		Wg26SN:    fmt.Sprintf("%d%d", start, end),
	}
}

// ParseFromWg26Number 从维根26标准字面卡号解析
func ParseFromWg26Number(wg26Number string) *Wg26Id {
	std := fmt.Sprintf("%8s", wg26Number)
	bytes := make([]byte, 2)
	bigOrder.PutUint16(bytes, uint16(ToInt64(std[3:])))
	return ParseFromWg26([3]byte{
		byte(ToInt64(strings.TrimLeft(std[:3], " "))),
		bytes[0],
		bytes[1]})
}
