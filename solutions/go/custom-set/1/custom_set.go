package stringset

import (
    "strconv"
    "strings"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set struct {
    seen		map[string]struct{}
    members		[]string
}

func New() Set {
	return Set{
        seen: make(map[string]struct{}),
        members: make([]string, 0),
    }
}

func NewFromSlice(l []string) Set {
	s := New()
    for _, element := range l {
        s.Add(element)
    }
    return s
}

func (s Set) String() string {
    quoted := make([]string, len(s.members))
    for i, member := range s.members {
        quoted[i] = strconv.Quote(member)
    }
    return "{" + strings.Join(quoted, ", ") + "}"
}

func (s Set) IsEmpty() bool {
	return len(s.members) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s.seen[elem]
    return ok
}

func (s *Set) Add(elem string) {
	if _, ok := s.seen[elem]; !ok {
        s.seen[elem] = struct{}{}
        s.members = append(s.members, elem)
    }
}

func Subset(s1, s2 Set) bool {
	if s1.IsEmpty() {
        return true
    }

    for _, member := range s1.members {
        if !s2.Has(member) {
            return false
        }
    }

    return true
}

func Disjoint(s1, s2 Set) bool {
	for _, member := range s1.members {
        if s2.Has(member) {
            return false
        }
    }

    return true
}

func Equal(s1, s2 Set) bool {
	if len(s1.members) != len(s2.members) {
        return false
    }

    for _, member := range s1.members {
        if !s2.Has(member) {
            return false
        }
    }

    return true
}

func Intersection(s1, s2 Set) Set {
	newSet := New()
    if len(s2.members) < len(s1.members) {
        s1, s2 = s2, s1
    }

    for _, member := range s1.members {
        if s2.Has(member) {
            newSet.Add(member)
        }
    }

    return newSet
}

func Difference(s1, s2 Set) Set {
	newSet := New()

    for _, member := range s1.members {
        if !s2.Has(member) {
            newSet.Add(member)
        }
    }

    return newSet
}

func Union(s1, s2 Set) Set {
	newSet := New()
    for _, member := range s1.members {
        newSet.Add(member)
    }
    for _, member := range s2.members {
        newSet.Add(member)
    }

    return newSet
}
