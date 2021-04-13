package mygoutil

import "testing"

func TestRandStrByCharset(t *testing.T) {
	t.Error(RandStrByCharset(10, "abcasdfsadf"))
}

func TestRandStr(t *testing.T) {
	t.Error(RandStr())
}