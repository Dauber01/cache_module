package cache

import "log"

// 缓存的行为
type Cache interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Del(string) error
	GetStat() Stat
}

// 创建新的缓存
func New(typ string) Cache {
	var c Cache
	if typ == "inmemory" {
		c = newInMemoryCache()
	}
	if c == nil {
		panic("unknow cache type " + typ)
	}
	log.Println(typ, "ready to serve")
	return c
}
