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
		t.Log("CardSN: ", id.CardSN)
		t.Log("CardHex: ", id.CardHex)
		t.Logf("Wg26Start: %d", id.Wg26Start)
		t.Logf("Wg26End: %d", id.Wg26End)
		t.Log("Wg26SN: ", id.Wg26SN)
		t.Log("FormatWg26SN: ", id.FormatWg26SN())

		if "0005653307" != id.CardSN {
			t.Fail()
		}
		if "56433B" != id.CardHex {
			t.Fail()
		}
		if 86 != id.Wg26Start {
			t.Fail()
		}
		if 17211 != id.Wg26End {
			t.Fail()
		}
		if "086,17211" != id.FormatWg26SN() {
			t.Fail()
		}
	}

	testFields(ParseFromCardNumber("0005653307"))
	testFields(ParseFromWg26([3]byte{0x56, 0x43, 0x3B}))
}

func TestParse1(t *testing.T) {
	id := ParseFromCardNumberValue(3659533)
	t.Log("CardSN: ", id.CardSN)
	t.Log("CardHex: ", id.CardHex)
	t.Logf("Wg26Start: %d", id.Wg26Start)
	t.Logf("Wg26End: %d", id.Wg26End)
	t.Log("Wg26SN: ", id.Wg26SN)
	t.Log("FormatWg26SN(): ", id.FormatWg26SN())

	if "0003659533" != id.CardSN {
		t.Fail()
	}
	if "37D70D" != id.CardHex {
		t.Fail()
	}
	if 55 != id.Wg26Start {
		t.Fail()
	}
	if 55053 != id.Wg26End {
		t.Fail()
	}
	if "055,55053" != id.FormatWg26SN() {
		t.Fail()
	}
}

func TestParse2(t *testing.T) {
	data := []byte{0xfb, 0x7c, 0x83, 0x00}
	wg26Id := fmt.Sprintf("%d", binary.LittleEndian.Uint32(data))
	id := ParseFromWg26Number(wg26Id)
	t.Log("CardSN: ", id.CardSN)
	t.Log("CardHex: ", id.CardHex)
	t.Logf("Wg26Start: %d", id.Wg26Start)
	t.Logf("Wg26End: %d", id.Wg26End)
	t.Log("Wg26SN: ", id.Wg26SN)
	t.Log("FormatWg26SN: ", id.FormatWg26SN())

	if "0005653307" != id.CardSN {
		t.Fail()
	}
	if "56433B" != id.CardHex {
		t.Fail()
	}
	if 86 != id.Wg26Start {
		t.Fail()
	}
	if 17211 != id.Wg26End {
		t.Fail()
	}
	if "086,17211" != id.FormatWg26SN() {
		t.Fail()
	}
}
