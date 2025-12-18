package encode

import (
    "strconv"
    "strings"
    "unicode"
)

func RunLengthEncode(input string) string {
    if input == "" {
        return ""
    }
	var sb strings.Builder

    count := 1
    runes := []rune(input)
    for i := 1; i < len(runes); i++ {
        if runes[i] == runes[i-1] {
            count++
        } else {
            if count > 1 {
                sb.WriteString(strconv.Itoa(count))
            }
            sb.WriteRune(runes[i-1])
            count = 1
        }
    }

    if count > 1 {
        sb.WriteString(strconv.Itoa(count))
    }
    sb.WriteRune(runes[len(runes)-1])

    return sb.String()
}

func RunLengthDecode(input string) string {
	var sb strings.Builder
	intStr := ""

    for _, r := range input {
        if unicode.IsDigit(r) {
            intStr += string(r)
        } else {
            count := 1
            if intStr != "" {
                count, _ = strconv.Atoi(intStr)
            }
            sb.WriteString(strings.Repeat(string(r), count))
            intStr = ""
        }
    }
    return sb.String()
}
