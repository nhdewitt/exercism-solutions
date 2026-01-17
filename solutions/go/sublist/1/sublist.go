package sublist

import "slices"

// Relation type is defined in relations.go file.

func isSublist(big, small []int) bool {
    if len(small) == 0 { return true }
    if len(small) > len(big) { return false }

    for j := 0; j <= len(big)-len(small); j++ {
        found := true
        if small[0] == big[j] {
            sub := big[j:j+len(small)]
            for k := range small {
                if small[k] != sub[k] {
                    found = false
                    break
                }
            }
            if found { return true }
        }
    }

    return false
}

func Sublist(l1, l2 []int) Relation {
	if slices.Equal(l1, l2) { return RelationEqual }
    if isSublist(l1, l2) { return RelationSuperlist }
    if isSublist(l2, l1) { return RelationSublist }
    return RelationUnequal
}
