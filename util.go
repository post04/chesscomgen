package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func getRandomGenProxy() string {
	proxy := config.Proxy
	if strings.Contains(proxy, "proxiware") || strings.Contains(proxy, "smartproxy") {
		return strings.ReplaceAll(proxy, "{rand}", fmt.Sprint(2001+rand.Intn(41500)))
	}
	return strings.ReplaceAll(proxy, "{rand}", RandStringBytes(10+rand.Intn(20)))
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const letterBytesLowercase = "abcdefghijklmnopqrstuvwxyz0123456789"

// RandStringBytes generates a random string x letters long
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func getUsername() string {
	if config.Username == "random" {
		return usernames[rand.Intn(len(usernames))]
	}
	return config.Username
}

func getRandomEmail() string {
	return RandStringBytes(10+rand.Intn(20)) + "@hotmail.com"
}
