// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
    "regexp"
    "strings"
    "unicode"
)

var responses = []string{
    "Sure.", "Whoa, chill out!", "Calm down, I know what I'm doing!", "Fine. Be that way!", "Whatever.",
}

func checkAllUpper(s string) bool {
    letterSeen := false
    for _, r := range s {
        if unicode.IsLetter(r) {
            letterSeen = true
        }
        if unicode.IsLetter(r) && unicode.ToUpper(r) != r {
            return false
        }
    }
    return letterSeen
}

func checkSilence(s string) bool {
    re := regexp.MustCompile("[^a-zA-z0-9]+")
    return re.ReplaceAllString(s, "") == ""
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
    var isQuestion, isAllCaps, isSilence bool
    remark = strings.TrimSpace(remark)

    if strings.HasSuffix(remark, "?") {
        isQuestion = true
    }
    if checkAllUpper(remark) {
        isAllCaps = true
    }
    if checkSilence(remark) {
        isSilence = true
    }

    switch {
    case isAllCaps && isQuestion:
        return responses[2]
    case isAllCaps:
        return responses[1]
    case isQuestion:
        return responses[0]
    case isSilence:
        return responses[3]
    default:
        return responses[4]
    }
}