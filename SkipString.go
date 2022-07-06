package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

type SkipString struct {
	skipChar  int
	skipIndex int
	str       string
}

func NewSkipString(pos int, str string) SkipString {
	lowercase := strings.ToLower(str)
	return SkipString{
		skipChar:  pos,
		str:       lowercase,
		skipIndex: 0,
	}
}

func (s SkipString) GetValueAsRuneSlice() []rune {
	return []rune(s.str)
}

func (s *SkipString) TransformRune(pos int) {
	runes := []rune(s.str)
	if regexp.MustCompile(`[a-zA-Z0-9]`).MatchString(string(runes[pos])) {
		if (s.skipIndex+1)%s.skipChar == 0 {
			uppercase := unicode.ToUpper(runes[pos])
			runes[pos] = uppercase
			s.str = string(runes)
		}
		s.skipIndex++
	}
}

func (s SkipString) String() string {
	return fmt.Sprintf("String with every %v characters capitalized: %v", s.skipChar, s.skipIndex)
}
