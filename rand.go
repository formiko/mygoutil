package mygoutil

import (
	"math/rand"
	"strings"
	"time"
)

func RandStrByCharset(length int, charset string) string {
	rand.Seed(time.Now().Unix())
	mp := make(map[int]string)
	i := 0
	for _, v := range charset {
		mp[i] = string(v)
		i++
	}
	var builder strings.Builder
	for i := 0; i < length; i++ {
		builder.WriteString(string(mp[rand.Intn(len(mp))]))
	}
	return builder.String()
}

func RandStr() string {
	return RandStrByCharset(16, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}