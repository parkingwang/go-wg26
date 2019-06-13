package wg26

import (
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
		if "086,17211" != id.Std() {
			t.Error("Not match")
		}
	}

	testFields(ParseFromSN("0005653307"))
	testFields(ParseFromWg26([3]byte{0x56, 0x43, 0x3B}))
}
