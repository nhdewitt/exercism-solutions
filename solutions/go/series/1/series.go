package series

func All(n int, s string) []string {
	slices := make([]string, 0)

    for i := range len(s)-n+1 {
        slices = append(slices, s[i:i+n])
    }
    return slices
}

func UnsafeFirst(n int, s string) string {
	return All(n, s)[0]
}
