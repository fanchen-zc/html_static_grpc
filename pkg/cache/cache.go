package cache

import (
	"github.com/coocood/freecache"
)

func NewCache(size int) *freecache.Cache {
	if size <= 0 {
		size = 10
	}
	cacheSize := size * 1024 * 1024
	return freecache.NewCache(cacheSize)
}
