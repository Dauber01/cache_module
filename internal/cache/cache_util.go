package cache

// 缓存状态
type Stat struct {
	Count     int64
	KeySize   int64
	ValueSize int64
}

// 新增缓存
func (s *Stat) add(k string, v []byte) {
	s.Count++
	s.KeySize += int64(len(k))
	s.ValueSize += int64(len(v))
}

// 减少缓存
func (s *Stat) del(k string, v []byte) {
	s.Count -= 1
	s.KeySize -= int64(len(k))
	s.ValueSize -= int64(len(v))
}
