package utils

import (
	"fmt"
	"net/url"
	"testing"
)

func TestHttpGet(t *testing.T) {
	u, _ := url.Parse("socks5://rola:123456@gate3.rola.vip:2059")
	u.User = url.UserPassword("rola", "123456")
	fmt.Println(u.User)
}
