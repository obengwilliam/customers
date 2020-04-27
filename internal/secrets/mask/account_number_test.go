// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package mask

import (
	"testing"
)

func TestAccountNumber(t *testing.T) {
	if v := AccountNumber(""); v != "**" {
		t.Errorf("got %q", v)
	}
	if v := AccountNumber("12"); v != "**" {
		t.Errorf("got %q", v)
	}
	if v := AccountNumber("123"); v != "*23" {
		t.Errorf("got %q", v)
	}
	if v := AccountNumber("1234"); v != "**34" {
		t.Errorf("got %q", v)
	}
	if v := AccountNumber("123456789"); v != "*******89" {
		t.Errorf("got %q", v)
	}
}
