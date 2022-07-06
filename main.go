package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/mapper_string/mapper"
)

const expected = "asPirAtiOn.cOm"

func main() {
	//Part 1
	capitalized := CapitalizeEveryThirdAlphanumericChar("Aspiration.Com")
	fmt.Println("capitalized: " + capitalized)
	if capitalized != expected {
		panic("wrong output!")
	}

	//Part 2
	s := NewSkipString(3, "Aspiration.com")
	mapper.MapString(&s)
	fmt.Println(s)

}

func CapitalizeEveryThirdAlphanumericChar(s string) string {
	lowercase := strings.ToLower(s)

	//strings are immutable in go, so create rune array
	runes := []rune(lowercase)

	// keep manual track of index, as we will be ignoring non-alphanumeric chars
	index := 0
	for i, char := range runes {
		if regexp.MustCompile(`[a-zA-Z0-9]`).MatchString(string(char)) {
			if (index+1)%3 == 0 {
				upper := unicode.ToUpper(runes[i])
				runes[i] = upper
			}
			index++
		}
	}

	return string(runes)
}
