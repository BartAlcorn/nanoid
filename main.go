package nanoid

/*
	Initially copied from: https://github.com/aidarkhanov/nanoid/blob/master/nanoid.go
	License: https://github.com/aidarkhanov/nanoid/blob/master/LICENSE (MIT license)
	* Any errors due to changes or modifications to the original code is on me.

	A tiny and fast Go unique string generator

	Safe. It uses cryptographically strong random APIs and tests distribution of symbols.
	Compact. It uses a larger alphabet than UUID (A-Za-z0-9_-). So ID size was reduced from 36 to 21 symbols.
*/

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/bits"
	"strings"
)

const (
	// DefaultAlphabet is the default alphabet for Nano ID.
	DefaultAlphabet = "-_0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	// Only Alpha. Used as the first charactor of WebSafe IDs,
	// as numbers as not allowed as the first character of an HTML element ID
	AlphaOnly = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	// No Punctuation, but all upper and lower alpha and all numeric characters
	AlphaNumeric = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	// DefaultSize is the default size for Nano ID.
	DefaultSize = 21
)

// BytesGenerator represents random bytes buffer.
type BytesGenerator func(step int) ([]byte, error)

// generateRandomBuffer generates a random buffer of the given step size.
// It takes an integer step as a parameter and returns a byte slice and an error.
func generateRandomBuffer(step int) ([]byte, error) {
	buffer := make([]byte, step)
	if _, err := rand.Read(buffer); err != nil {
		return nil, err
	}
	return buffer, nil
}

// FormatString generates a random string based on BytesGenerator, alphabet and size.
func FormatString(generateRandomBuffer BytesGenerator, alphabet string, size int) (string, error) {
	mask := 2<<uint32(31-bits.LeadingZeros32(uint32(len(alphabet)-1|1))) - 1
	step := int(math.Ceil(1.6 * float64(mask*size) / float64(len(alphabet))))

	id := new(strings.Builder)
	id.Grow(size)

	for {
		randomBuffer, err := generateRandomBuffer(step)
		if err != nil {
			return "", err
		}

		for i := 0; i < step; i++ {
			currentIndex := int(randomBuffer[i]) & mask

			if currentIndex < len(alphabet) {
				if err := id.WriteByte(alphabet[currentIndex]); err != nil {
					return "", err
				} else if id.Len() == size {
					return id.String(), nil
				}
			}
		}
	}
}

// GenerateString generates a random string based on alphabet and size.
func GenerateString(alphabet string, size int) (string, error) {
	id, err := FormatString(generateRandomBuffer, alphabet, size)
	if err != nil {
		return "", err
	}
	return id, nil
}

// New generates a random string.
func New() (string, error) {
	var sb strings.Builder

	//ensure we don't start with a number or punctuation
	clipped := strings.TrimLeft(DefaultAlphabet, "-_0123456789")
	str, err := GenerateString(clipped, 1)
	if err != nil {
		fmt.Printf("%v %v\n", "ERROR creating NanoID", err)
	}

	//generate the remaining characters
	sb.WriteString(str)
	str, err = GenerateString(DefaultAlphabet, DefaultSize-1)

	sb.WriteString(str)
	return sb.String(), err
}

// for times when returning an error is not desired
func NewMust() string {
	var str string
	var err error
	str, err = New()
	if err != nil {
		fmt.Printf("%v %v\n", "ERROR creating NanoID", err)
	}
	return str
}

// New generates a random string, but only 14 characters long.
func New14() (string, error) {
	return GenerateString(DefaultAlphabet, 14)
}

// WebSafeID, returns a 16 character (4x4, dash separated) string
// that is safe to use as an HTML ID
func WebSafeID() string {
	var p = make([]string, 4)
	for i := 0; i < 4; i++ {
		var s string
		str, err := GenerateString(AlphaOnly, 1)
		if err != nil {
			fmt.Printf("%v %v\n", "ERROR creating NanoID", err)
		}
		s = str
		for j := 0; j < 1; j++ {
			str, err = GenerateString(AlphaNumeric, 3)
			if err != nil {
				fmt.Printf("%v %v\n", "ERROR creating NanoID", err)
			}
			s += str
		}
		p[i] = s
	}
	return strings.Join(p, "-")
}
