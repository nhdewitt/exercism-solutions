package thefarm

import "errors"
import "fmt"

func DivideFood(f FodderCalculator, numberOfCows int) (float64, error) {
    fodderAmt := 0.0
    fodderAmt, err := f.FodderAmount(numberOfCows)
    if err != nil {
        return 0.0, err
    } 

    factor, err := f.FatteningFactor()
    if err != nil {
        return 0.0, err
    }

    return float64(fodderAmt * factor) / float64(numberOfCows), nil
}

func ValidateInputAndDivideFood(f FodderCalculator, numberOfCows int) (float64, error) {
    if numberOfCows > 0 {
        dividedFood, err := DivideFood(f, numberOfCows)
        if err != nil {
            return 0.0, err
        }
        return dividedFood, nil
    }

    return 0.0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
    InvalidInt		int
    Message			string
}

func (e *InvalidCowsError) Error() string {
    return e.Message
} 

func ValidateNumberOfCows(n int) error {
    switch {
    case n < 0:
        return &InvalidCowsError{
            InvalidInt: n,
            Message: fmt.Sprintf("%d cows are invalid: there are no negative cows", n)}
    case n == 0:
        return &InvalidCowsError{
            InvalidInt: n,
            Message: fmt.Sprintf("0 cows are invalid: no cows don't need food")}
    default:
        return nil
    }
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
