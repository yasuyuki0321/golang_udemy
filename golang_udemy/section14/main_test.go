package main

// テストファイルの名前
// パッケージ名 + _test.go

import "testing"

var Debug bool = false

// Test関数名はテスト用の関数
func TestIsOne(t *testing.T) {
	i := 1
	if Debug {
		t.Skip("スキップさせる")
	}

	v := IsOne(i)
	if !v {
		t.Errorf("%v != %v", i, 1)
	}
}
