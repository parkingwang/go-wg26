package wg26

import (
	"encoding/binary"
	"fmt"
	"testing"
)

//
// Author: 陈哈哈 bitschen@163.com
//

func TestParseFromSN(t *testing.T) {
	testFields := func(id *Wg26Id) {
		if "0005653307" != id.Number {
			t.Error("Not match")
		}
		if "56433B" != id.Wg26Hex {
			t.Error("Not match")
		}
		if "56433B" != fmt.Sprintf("%X", id.Wg26Bytes) {
			t.Error("Not match")
		}
		if 86 != id.Wg26Start {
			t.Error("not match")
		}
		if 17211 != id.Wg26End {
			t.Error("not match")
		}
		if "086,17211" != id.Std() {
			t.Error("Not match")
		}
	}

	testFields(ParseFromCardNumber("0005653307"))
	testFields(ParseFromWg26([3]byte{0x56, 0x43, 0x3B}))
}

func TestParse1(t *testing.T) {
	id := ParseFromCardNumberValue(3659533)
	t.Log("Number: ", id.Number)
	t.Log("Wg26Hex: ", id.Wg26Hex)
	t.Logf("Wg26Bytes: %X", id.Wg26Bytes)
	t.Logf("Wg26Start: %d", id.Wg26Start)
	t.Logf("Wg26End: %d", id.Wg26End)
	t.Log("Std: ", id.Std())
}

func TestParse2(t *testing.T) {
	data := []byte{0xfb, 0x7c, 0x83, 0x00}
	wg26Id := fmt.Sprintf("%d", binary.LittleEndian.Uint32(data))
	id := ParseFromWg26Number(wg26Id)
	t.Log("Number: ", id.Number)
	t.Log("Wg26Hex: ", id.Wg26Hex)
	t.Logf("Wg26Bytes: %X", id.Wg26Bytes)
	t.Logf("Wg26Start: %d", id.Wg26Start)
	t.Logf("Wg26End: %d", id.Wg26End)
	t.Log("Std: ", id.Std())
}
