package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	acc := initial
    if s.Length() == 0 {
        return initial
    }

    for _, i := range s {
        acc = fn(acc, i)
    }

    return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	acc := initial
	length := s.Length()
    if length == 0 {
        return initial
    }

    for i := length-1; i >= 0; i-- {
        acc = fn(s[i], acc)
    }

    return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	new := make(IntList, 0)

    for _, i := range s {
        if fn(i) {
            new = append(new, i)
        }
    }

    return new
}

func (s IntList) Length() int {
	count := 0

    for range s {
        count++
    }

    return count
}

func (s IntList) Map(fn func(int) int) IntList {
	new := make(IntList, 0)

    for _, i := range s {
        new = append(new, fn(i))
    }

    return new
}

func (s IntList) Reverse() IntList {
    length := s.Length()
	new := make(IntList, length)
    for i, j := 0, length-1; i < length; i, j = i+1, j-1 {
        new[i] = s[j]
    }

    return new
}

func (s IntList) Append(lst IntList) IntList {
    origLen := s.Length()
    appLen := lst.Length()
    new := make(IntList, origLen+appLen)

    for i := 0; i < origLen; i++ {
        new[i] = s[i]
    }
    j := len(s)
    for i := 0; i < appLen; i++ {
        new[i+j] = lst[i]
    }

    return new
}

func (s IntList) Concat(lists []IntList) IntList {
	new := s

    for _, list := range lists {
        new = new.Append(list)
    }

    return new
}
