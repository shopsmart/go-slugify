package slugify

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mozillazg/go-unidecode"
)

// Separator separator between words
var Separator = "-"

// StripChar removes invalid slug characters
var StripChar = regexp.MustCompile(fmt.Sprint("['\"`]", ""))

// ReInValidChar match invalid slug string
var ReInValidChar = regexp.MustCompile(fmt.Sprint("[^a-zA-Z0-9]+"))

// Version return version
func Version() string {
	return "0.2.4"
}

// Slugify implements make a pretty slug from the given text.
// e.g. Slugify("kožušček hello world") => "kozuscek-hello-world"
func Slugify(s string) string {
	return slugify(s)
}

func slugify(s string) string {
	s = unidecode.Unidecode(s)
	s = stripChar(s)
	s = replaceInValidCharacter(s, Separator)
	s = strings.Trim(s, Separator)
	s = strings.ToLower(s)
	return s
}

func stripChar(s string) string {
	s = StripChar.ReplaceAllString(s, "")
	return s
}

func replaceInValidCharacter(s, repl string) string {
	s = ReInValidChar.ReplaceAllString(s, repl)
	return s
}
