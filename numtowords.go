package numtowords

import "fmt"

// MaxNumber is the maximum number that can be converted to words.
const MaxNumber = 9999

// Convert converts a number to its word representation.
func Convert(number int) string {
	text := ""
	if number < 0 || number > MaxNumber {
		return "Number out of range"
	}

	units := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	unitsTeens := []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	unitsTens := []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

	if number >= 1000 {
		numThu := number / 1000
		text += fmt.Sprintf("%s thousand ", units[numThu])
		number = number - (numThu * 1000)
	}
	// thplace := procNumber(&number, 1000, units)
	// if thplace != "" {
	// 	text += fmt.Sprintf("%s thousand ", thplace)
	// }

	if number >= 100 {
		numHun := number / 100
		text += fmt.Sprintf("%s hundred ", units[numHun])
		number = number - (numHun * 100)
	}
	// fmt.Println("Arun", &number)
	// fmt.Println("Arun1", number)
	// hplace := procNumber(&number, 100, units)
	// if hplace != "" {
	// 	text += fmt.Sprintf("%s hundred ", hplace)
	// }

	if number >= 20 {
		numTen := number / 10
		text += fmt.Sprintf("%s ", unitsTens[numTen])
		number = number - (numTen * 10)
	}
	// tnplace := procNumber(&number, 20, unitsTens)
	// if tnplace != "" {
	// 	text += fmt.Sprintf("%s ", tnplace)
	// }
	if number >= 10 {
		text += fmt.Sprintf("%s ", unitsTeens[number-10])
		number = 0
	}

	// tplace := procNumber(&number, 10, unitsTeens)
	// if tplace != "" {
	// 	text += fmt.Sprintf("%s ", tplace)
	// }

	if number > 0 {
		text += fmt.Sprintf("%s", units[number])
	}

	// oplace := procNumber(&number, 1, units)
	// if oplace != "" {
	// 	text += fmt.Sprintf("%s ", oplace)
	// }
	return text
}

func procNumber(number *int, div int, myunits []string) string {
	newNum := *number / div
	if *number >= 10 && *number < 20 {
		newNum = *number - 10
		div = 1
	}
	fmt.Println("number " + fmt.Sprint(*number) + " , newNumber " + fmt.Sprint(newNum))
	texts := fmt.Sprintf("%s", myunits[newNum])
	*number = *number - (newNum * div)
	return texts
}
