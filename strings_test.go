package mygoutil

import "testing"

func TestSortString(t *testing.T) {
	t.Logf(SortString("2312716669"))
	t.Logf(SortString("312陈312"))
	t.Logf(SortString("陈2312716669果"))
}