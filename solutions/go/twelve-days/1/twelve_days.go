package twelve

import "strings"

var gifts = []string{
    "a Partridge in a Pear Tree", "two Turtle Doves", "three French Hens", "four Calling Birds",
    "five Gold Rings", "six Geese-a-Laying", "seven Swans-a-Swimming", "eight Maids-a-Milking",
    "nine Ladies Dancing", "ten Lords-a-Leaping", "eleven Pipers Piping", "twelve Drummers Drumming",
}
var ordinals = []string{
    "first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth",
    "eleventh", "twelfth",
}

func Verse(i int) string {
    var sb strings.Builder
    sb.WriteString("On the " + ordinals[i-1] + " day of Christmas my true love gave to me: ")
    if i > 1 {    
    	for j := i - 1; j > 0; j-- {
            sb.WriteString(gifts[j] + ", ")
        }
        sb.WriteString("and ")
    }
    sb.WriteString(gifts[0] + ".")

    return sb.String()
}

func Song() string {
	output := make([]string, 0, 12)
    for i := range 12 {
        output = append(output, Verse(i+1))
    }

    return strings.Join(output, "\n")
}
