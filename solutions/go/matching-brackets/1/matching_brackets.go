package brackets

type Stack[T any] struct {
    items	[]T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if s.IsEmpty() {
        var zero T
        return zero, false
    }
    index := len(s.items) - 1
    item := s.items[index]
    s.items = s.items[:index]
    return item, true
}

func (s *Stack[T]) IsEmpty() bool {
    return len(s.items) == 0
}

func Bracket(input string) bool {
    bracketMap := map[rune]rune{
        '[': ']',
        '(': ')',
        '{': '}',
    }

    stack := Stack[rune]{}

    for _, r := range input {
        if r == '[' || r == '(' || r == '{' {
            stack.Push(r)
            continue
        }
        if r == ']' || r == ')' || r == '}' {
            top, ok := stack.Pop()
            if !ok { return false }
            if bracketMap[top] != r { return false }
        }
    }

    return stack.IsEmpty()
}
