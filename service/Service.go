package service

import (
	"github.com/coocood/freecache"
)

var cache *freecache.Cache

func Setup() {
	size := 20 * 1024 * 1024 // 缓存大小 20M
	cache = freecache.NewCache(size)

}

func GetCache(key []byte) string {
	got, err := cache.Get(key)
	if err != nil {
		return ""
	}
	return string(got)
}

func SetCache(key, val []byte, expire int) {
	cache.Set(key, val, expire)
}
