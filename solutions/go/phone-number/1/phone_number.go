package phonenumber

import (
    "fmt"
    "regexp"
)

var validPhone = regexp.MustCompile(
    `^\s*(?:\+?1[\s.\-]*)?\(?(?P<area>[2-9]\d{2})\)?[\s.\-]*(?P<exch>[2-9]\d{2})[\s.\-]*(?P<line>\d{4})\s*$`,
)

func Number(phoneNumber string) (string, error) {
	m := validPhone.FindStringSubmatch(phoneNumber)
    if m == nil {
        return "", fmt.Errorf("invalid phone format")
    }

    idx := make(map[string]int)
    for i, name := range validPhone.SubexpNames() {
        if name != "" {
            idx[name] = i
        }
    }

    return m[idx["area"]] + m[idx["exch"]] + m[idx["line"]], nil
}

func AreaCode(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    
    return num[:3], nil
}

func Format(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }

    return fmt.Sprintf("(%s) %s-%s", num[:3], num[3:6], num[6:]), nil
}
