package flatten

func Flatten(nested interface{}) []interface{} {
	flat := []interface{}{}

    switch n := nested.(type) {
    case []interface{}:
        for _, val := range n {
            if val != nil {
                flat = append(flat, Flatten(val)...)
            }
        }
    case interface{}:
        flat = append(flat, n)
    }

    return flat
}
