package flags

import (
	"flag"

	"github.com/hgochev/password-generator/internal/models"
)

func ParseArguments() models.Options {
	length := flag.Int("l", 16, "length of the generated password")
	noNumbers := flag.Bool("nn", false, "don't include numbers")
	noSpecialChars := flag.Bool("ns", false, "don't include special characters")
	noMixedCase := flag.Bool("nm", false, "don't use mixed case letters")

	flag.Parse()

	return models.Options{
		Length:         *length,
		NoNumbers:      *noNumbers,
		NoSpecialChars: *noSpecialChars,
		NoMixedCase:    *noMixedCase,
	}
}
