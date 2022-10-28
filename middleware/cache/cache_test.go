package cache

import (
	"testing"
	"time"

	"github.com/fanjindong/go-cache"
)

func mainxxx(t *testing.T) {
	c := cache.NewMemCache()
	c.Set("a", 1)
	c.Set("b", 1, cache.WithEx(1*time.Second))
	time.Sleep(1 * time.Second)
	c.Get("a") // 1, true
	c.Get("b") // nil, false
}
