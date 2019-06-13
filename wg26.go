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

// 维根26编码
type Wg26Id struct {
	Number    string  // 原始十位卡号
	Wg26Hex   string  // 维根26十六进制
	Wg26Bytes [3]byte // 维根26字节码
	Wg26Start uint16  // 维根26，前段
	Wg26End   uint16  // 维根26，后段
}

// Std 返回维根26标准卡号格式化字符串
func (id *Wg26Id) Std() string {
	return "0" + fmt.Sprintf("%d,%d", id.Wg26Start, id.Wg26End)
}

// 从原始卡号解析
func ParseFromSN(number string) *Wg26Id {
	nInt := TrimZero(number)
	nHex := fmt.Sprintf("%06X", nInt)
	bytes, _ := hex.DecodeString(nHex)
	return &Wg26Id{
		Number:    number,
		Wg26Hex:   nHex,
		Wg26Bytes: [3]byte{bytes[0], bytes[1], bytes[2]},
		Wg26Start: uint16(bytes[0]),
		Wg26End:   binary.BigEndian.Uint16(bytes[1:]),
	}
}

// 从维根26国际标准编码解析
func ParseFromWg26(wg26std [3]byte) *Wg26Id {
	num := fmt.Sprintf("%d", binary.BigEndian.Uint32([]byte{0, wg26std[0], wg26std[1], wg26std[2]}))
	return &Wg26Id{
		Number:    AppendZero(num),
		Wg26Hex:   fmt.Sprintf("%06X", wg26std),
		Wg26Bytes: wg26std,
		Wg26Start: uint16(wg26std[0]),
		Wg26End:   binary.BigEndian.Uint16(wg26std[1:]),
	}
}

// 从Uint32类型的维根26国际标准编码解析
func ParseFromUint32(id uint32) *Wg26Id {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, id)
	return ParseFromWg26([3]byte{b[1], b[2], b[3]})
}

func TrimZero(card string) int64 {
	card = strings.TrimLeft(card, "0")
	v, _ := strconv.ParseInt(card, 10, 64)
	return v
}

func AppendZero(card string) string {
	s := 10 - len(card)
	zeros := ""
	for i := 0; i < s; i++ {
		zeros += "0"
	}
	return zeros + card
}
