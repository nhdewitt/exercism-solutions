package cipher

import (
    "crypto/rand"
    "math/big"
    "strings"
)

const asciiLowercase	= "abcdefghijklmnopqrstuvwxyz"

type shift struct {
    shiftVal	int
}

type vigenere struct {
    shiftKey	string
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance == 0 || distance > 25 {
        return nil
    }
	return shift{shiftVal: distance}
}

func (c shift) Encode(input string) string {
	return shiftText(input, c.shiftVal)
}

func (c shift) Decode(input string) string {
	return shiftText(input, -c.shiftVal)
}

func NewVigenere(key string) Cipher {
    if !isValidKey(key) {
        return nil
    }
    
	if key == "" {
        key = generateRandomKey(100)
    }

    return vigenere{shiftKey: key}
}

func (v vigenere) Encode(input string) string {
	return v.transform(input, 1)
}

func (v vigenere) Decode(input string) string {
	return v.transform(input, -1)
}

func (v vigenere) transform(input string, direction int) string {
    var result strings.Builder
    keyPos := 0

    for _, r := range strings.ToLower(input) {
        if r >= 'a' && r <= 'z' {
            shift := int(v.shiftKey[keyPos%len(v.shiftKey)] - 'a')
            transformed := (r - 'a' + rune(direction*shift) + 26) % 26 + 'a'
            result.WriteRune(transformed)
            keyPos++
        }
    }
    return result.String()
}

func shiftText(input string, distance int) string {
    var result strings.Builder
    normalizedShift := ((distance % 26) + 26) % 26

    for _, r := range strings.ToLower(input) {
        if r >= 'a' && r <= 'z' {
            shifted := (r - 'a' + rune(normalizedShift)) % 26 + 'a'
            result.WriteRune(shifted)
        }
    }
    return result.String()
}

func isValidKey(key string) bool {
    if key == "" {
        return false
    }

    hasNonA := false
    for _, r := range key {
        if r < 'a' || r > 'z' {
            return false
        }
        if r != 'a' {
            hasNonA = true
        }
    }

    return hasNonA
}

func generateRandomKey(length int) string {
    var key strings.Builder
    for _ = range length {
        n, _ := rand.Int(rand.Reader, big.NewInt(26))
        key.WriteByte(asciiLowercase[n.Int64()])
    }
    return key.String()
}