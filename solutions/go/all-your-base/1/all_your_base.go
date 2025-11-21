package allyourbase

import (
    "fmt"
    "slices"
)

func intPow(base, exp int) int {
    prod := 1
    for exp > 0 {
        if exp%2 == 1 {
            prod *= base
        }
        base *= base
        exp /= 2
    }
    return prod
}

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
        return nil, fmt.Errorf("input base must be >= 2")
    }
    if outputBase < 2 {
        return nil, fmt.Errorf("output base must be >= 2")
    }
    var nonZeroSeen bool
    for _, d := range inputDigits {
        if d < 0 || d >= inputBase {
            return nil, fmt.Errorf("all digits must satisfy 0 <= d < input base")
        }
        if d != 0 {
            nonZeroSeen = true
        }
    }

    if !nonZeroSeen {
        return []int{0}, nil
    }

    sum := 0
    res := make([]int, 0)

    position := 0
    for i := len(inputDigits)-1; i >= 0; i-- {
        sum += inputDigits[i] * intPow(inputBase, position)
        position++
    }
    if sum == 0 {
        return []int{0}, nil
    }
    for sum > 0 {
        res = append(res, sum % outputBase)
        sum /= outputBase
    }

    slices.Reverse(res)
    return res, nil
}
