package gen

type Stack[T any] struct {
	stack []T
}

func (s *Stack[T]) Length() int {
	return len(s.stack)
}

func (s *Stack[T]) Push(item T) {
	s.stack = append(s.stack, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.Length() == 0 {
		var zero T
		return zero, false
	}
	idx := len(s.stack) - 1
	item := s.stack[idx]
	s.stack = s.stack[:idx]
	return item, true
}
