package day19

type Set struct {
	m map[location]struct{}
}

//MakeSet initialize the set
func MakeSet() *Set {
	return &Set{
		m: make(map[location]struct{}),
	}
}

func (s *Set) Contains(key location) bool {
	_, exists := (*s).m[key]
	return exists
}

func (s *Set) Add(key location) {
	(*s).m[key] = struct{}{}
}

func (s *Set) Remove(key location) {
	_, exists := (*s).m[key]
	if exists {
		delete((*s).m, key)
	}
}

func (s *Set) Size() int {
	return len((*s).m)
}
