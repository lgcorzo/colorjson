// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the https://golang.org/LICENSE file.

package colorjson

import (
	"testing"
)

type sampleStruct struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func TestMarshalAndUnmarshal(t *testing.T) {
	s := sampleStruct{
		Name:  "test",
		Value: 42,
	}

	b, err := Marshal(s)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	if len(b) == 0 {
		t.Fatalf("Marshal returned empty slice")
	}

	var res sampleStruct
	err = Unmarshal([]byte(`{"name":"test","value":42}`), &res)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if res.Name != "test" || res.Value != 42 {
		t.Fatalf("Unexpected unmarshal output: %+v", res)
	}
}
