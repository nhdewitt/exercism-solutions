package allergies

import "math/bits"

var allergens = []struct{
    name string
    score uint
}{
    {"eggs", 1},
    {"peanuts", 2},
    {"shellfish", 4},
    {"strawberries", 8},
    {"tomatoes", 16},
    {"chocolate", 32},
    {"pollen", 64},
    {"cats", 128},
}

func Allergies(allergies uint) []string {
	res := make([]string, 0, bits.OnesCount(allergies))

    for _, allergen := range allergens {
        if AllergicTo(allergies, allergen.name) {
            res = append(res, allergen.name)
        }
    }

    return res
}

func AllergicTo(allergies uint, allergen string) bool {
	for _, a := range allergens {
        if a.name == allergen {
            return allergies & a.score != 0
        }
    }
    return false
}
