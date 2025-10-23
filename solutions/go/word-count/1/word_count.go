package wordcount

import (
    "regexp"
    "strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
    re := regexp.MustCompile(`'\B|\B'`)
    clean := re.ReplaceAllString(phrase, "")
    re = regexp.MustCompile(`[^a-z0-9'\s]+`)
    clean = re.ReplaceAllString(strings.ToLower(clean), " ")
    words := []string{}
    re = regexp.MustCompile(`(^'|'$)`)
    for _, w := range strings.Fields(clean) {
        w = re.ReplaceAllString(w, "")
        if w != "" {
            words = append(words, w)
        }
    }
    count := make(Frequency)
    for _, word := range words {
        count[word]++
    }
    return count
}
