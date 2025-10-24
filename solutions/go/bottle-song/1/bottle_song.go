package bottlesong

import (
    "fmt"
    "strings"
    "unicode"
    "unicode/utf8"
)

func capitalize(s string) string {
    r, size := utf8.DecodeRuneInString(s)
    firstChar := string(unicode.ToUpper(r))
    rest := strings.ToLower(s[size:])

    return firstChar + rest
}

func pluralize(n int, s string) string {
    if n != 1 {
        return s + "s"
    }
    return s
}

func Recite(startBottles, takeDown int) []string {
        intMap := map[int]string{
        10: "ten",
        9: "nine",
        8: "eight",
        7: "seven",
        6: "six",
        5: "five",
        4: "four",
        3: "three",
        2: "two",
        1: "one",
        0: "no",
    }

    song := make([]string, 0)
    lastVerse := (startBottles - takeDown + 1)

    for i := startBottles; i > startBottles - takeDown; i-- {

        for range 2 {
            song = append(song, fmt.Sprintf("%s green %s hanging on the wall,", capitalize(intMap[i]), pluralize(i, "bottle")))
        }
        song = append(song, "And if one green bottle should accidentally fall,")
        song = append(song, fmt.Sprintf("There'll be %s green %s hanging on the wall.", intMap[i-1], pluralize(i-1, "bottle")))
		if i != 1 && i != lastVerse {
            song = append(song, "")
        }
    }

    return song
}
