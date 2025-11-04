package yacht

type Scores struct {
    Numbers [6]int
    FullHouse int
    FourOfAKind int
    LittleStraight int
    BigStraight int
    Choice int
    Yacht int
    DiceCount map[int]int
    DiceSum int
}

func (s *Scores) CountSingles() {
    for i := range 6 {
        s.Numbers[i] = (i + 1) * s.DiceCount[i+1]
    }
}

func (s *Scores) CheckFullHouse() {
    fh := make([]int, 2)
    for k, v := range s.DiceCount {
        if v == 3 {
            fh[1] = k
        } else if v == 2 {
            fh[0] = k
        }
    }

    if fh[0] > 0 && fh[1] > 0 {
        s.FullHouse = (2 * fh[0]) + (3 * fh[1])
    }
}

func (s *Scores) CheckFourOfAKind() {
    for k, v := range s.DiceCount {
        if v >= 4 {
            s.FourOfAKind = 4 * k
            return
        }
    }
}

func (s *Scores) CheckStraights() {
    if len(s.DiceCount) != 5 {
        return
    }
    if s.DiceSum == 15 {
        s.LittleStraight = 30
    } else if s.DiceSum == 20 {
        s.BigStraight = 30
    }
}

func (s *Scores) CheckYacht() {
    if len(s.DiceCount) == 1 {
        s.Yacht = 50
    }
}

func (s *Scores) CountDice(dice []int) {
    for _, die := range dice {
        s.DiceCount[die]++
        s.DiceSum += die
    }
}

func NewScores(dice []int) *Scores {
    s := &Scores{
        DiceCount:	make(map[int]int),
    }
    s.CountDice(dice)
    s.CountSingles()
    s.CheckFullHouse()
    s.CheckFourOfAKind()
    s.CheckStraights()
    s.CheckYacht()
    return s
}

func Score(dice []int, category string) int {
    s := NewScores(dice)

    switch category {
    case "ones":
        return s.Numbers[0]
    case "twos":
        return s.Numbers[1]
    case "threes":
        return s.Numbers[2]
    case "fours":
        return s.Numbers[3]
    case "fives":
        return s.Numbers[4]
    case "sixes":
        return s.Numbers[5]
    case "full house":
        return s.FullHouse
    case "four of a kind":
        return s.FourOfAKind
    case "little straight":
        return s.LittleStraight
    case "big straight":
        return s.BigStraight
    case "choice":
        return s.DiceSum
    case "yacht":
        return s.Yacht
    default:
        return 0
    }
}
