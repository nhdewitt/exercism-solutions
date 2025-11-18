package piglatin

import "strings"

func findVowelIdx(s string) int {
    vowels := "aeiouAEIOU"

    for i, r := range s {
        if strings.ContainsRune(vowels, r) {
            return i
        }
    }
    for i, r := range s {
        if strings.ContainsRune("y", r) {
            return i
        }
    }

    return -1
}

func Sentence(sentence string) string {
    fields := strings.Fields(sentence)
    res := make([]string, len(fields))
	prefixes := []string{"a","e","i","o","u","xr","yt"}

    for i, field := range fields {
        found := false
        for _, prefix := range prefixes {
            if strings.HasPrefix(field, prefix) {
                res[i] = field + "ay"
                found = true
                break
            }
        }
        if !found && strings.HasPrefix(field, "qu") {
            res[i] = field[2:] + "quay"
            found = true
        }
		if found { continue }

        idx := findVowelIdx(field)
        if idx == -1 {
            res[i] = ""
            continue
        }
        if idx == 2 && string(field[1]) == "q" {
            res[i] = field[idx+1:] + field[:idx+1] + "ay"
        } else {
            res[i] = field[idx:] + field[:idx] + "ay"
        }
    }

	return strings.Join(res, " ")
}
