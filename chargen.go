package stringen

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

var AlphaChars = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
	"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
	"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var NumericChars = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
var PunctuationChars = []string{".", ",", "?", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")"}

type CharType int

const (
	CharTypeAlpha = iota
	CharTypeAlphaNumeric
	CharTypeAlphaNumericSpecial
)

func (ct CharType) String() string {
	return [...]string{"alpha", "alphanum", "all"}[ct]
}

func GenRandomCharacters(numChars int, charType CharType) string {
	var randomStr string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numChars; i++ {
		switch charType {
		case CharTypeAlpha:
			randomIdx := rand.Intn(len(AlphaChars))
			randomStr = fmt.Sprintf("%s%s", randomStr, AlphaChars[randomIdx])
		case CharTypeAlphaNumeric:
			randomType := rand.Intn(2)
			switch randomType {
			case CharTypeAlpha:
				randomIdx := rand.Intn(len(AlphaChars))
				randomStr = fmt.Sprintf("%s%s", randomStr, AlphaChars[randomIdx])
			case CharTypeAlphaNumeric:
				randomIdx := rand.Intn(len(NumericChars))
				randomStr = fmt.Sprintf("%s%s", randomStr, NumericChars[randomIdx])
			}
		case CharTypeAlphaNumericSpecial:
			randomType := rand.Intn(3)
			switch randomType {
			case CharTypeAlpha:
				randomIdx := rand.Intn(len(AlphaChars))
				randomStr = fmt.Sprintf("%s%s", randomStr, AlphaChars[randomIdx])
			case CharTypeAlphaNumeric:
				randomIdx := rand.Intn(len(NumericChars))
				randomStr = fmt.Sprintf("%s%s", randomStr, NumericChars[randomIdx])
			case CharTypeAlphaNumericSpecial:
				randomIdx := rand.Intn(len(PunctuationChars))
				randomStr = fmt.Sprintf("%s%s", randomStr, PunctuationChars[randomIdx])
			}
		}
	}
	return randomStr
}

func GenSha256Hash(str string) string {
	hasher := sha256.New()
	hasher.Write([]byte(str))
	hash := hasher.Sum(nil)
	hashStr := hex.EncodeToString(hash)
	return hashStr
}
