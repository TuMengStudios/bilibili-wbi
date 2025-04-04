package wbi

import (
	"testing"
)

func TestWbiUrl(t *testing.T) {
	var url = "https://api.bilibili.com/x/space/wbi/arc/search?mid=object-mid"
	var newUrl, err = WbiSignURL(url)
	t.Log(newUrl, err)
}
