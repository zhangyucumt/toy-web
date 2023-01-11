package afterclass

type Set interface {
	Put(key string)
	Keys() []string
	Contains(key string) bool
	Remove(key string)
	// 如果之前已经有了，就返回旧的值，absent =false
	// 如果之前没有，就塞下去，返回 absent = true
	PutIfAbsent(key string) (old string, absent bool)
}

type SetStruct struct {
	container map[string]bool
}

func (s *SetStruct) Put(key string) {
	s.container[key] = true
}

func (s *SetStruct) Keys() []string {
	ret := make([]string, 0, len(s.container))
	for k := range s.container {
		ret = append(ret, k)
	}
	return ret
}

func (s *SetStruct) Contains(key string) bool {
	_, ok := s.container[key]
	return ok
}

func (s *SetStruct) Remove(key string) {
	delete(s.container, key)
}

func (s *SetStruct) PutIfAbsent(key string) (old string, absent bool) {
	if s.Contains(key) {
		old = key
		return
	} else {
		absent = true
		s.Put(key)
		return
	}
}
