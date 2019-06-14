package wg26

import (
	"testing"
)

//
// Author: 陈哈哈 bitschen@163.com
//

func TestIsDigits(t *testing.T) {
	if !IsDigits("0005653307") {
		t.Fail()
	}
	if IsDigits("ABC000") {
		t.Fail()
	}
}
