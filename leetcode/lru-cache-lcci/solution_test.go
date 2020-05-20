package lru_cache_lcci

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	t.Log(lru.Get(1))
	lru.Put(3, 3)
	lru.Get(2)
}
