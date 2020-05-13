package slugify

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mozillazg/go-unidecode"
)

// Separator separator between words
var Separator = "-"

// StripRepl replaces invalid characters
var StripRepl = ""

// SeparatorForRe for regexp
var SeparatorForRe = regexp.QuoteMeta(Separator)

// StripChar removes invalid slug characters
var StripChar = regexp.MustCompile(fmt.Sprintf("[%s'\"]", StripRepl))

// ReInValidChar match invalid slug string
var ReInValidChar = regexp.MustCompile(fmt.Sprintf("[^%sa-zA-Z0-9]", Separator))

// ReDupSeparatorChar match duplicate separator string
var ReDupSeparatorChar = regexp.MustCompile(fmt.Sprintf("%s{2,}", SeparatorForRe))

// Version return version
func Version() string {
	return "0.2.2"
}

// Slugify implements make a pretty slug from the given text.
// e.g. Slugify("kožušček hello world") => "kozuscek-hello-world"
func Slugify(s string) string {
	return slugify(s)
}

func slugify(s string) string {
	s = unidecode.Unidecode(s)
	s = stripChar(s, StripRepl)
	s = replaceInValidCharacter(s, Separator)
	s = removeDupSeparator(s)
	s = strings.Trim(s, Separator)
	s = strings.ToLower(s)
	return s
}

func stripChar(s, repl string) string {
	s = StripChar.ReplaceAllString(s, repl)
	return s
}

func replaceInValidCharacter(s, repl string) string {
	s = ReInValidChar.ReplaceAllString(s, repl)
	return s
}

func removeDupSeparator(s string) string {
	s = ReDupSeparatorChar.ReplaceAllString(s, Separator)
	return s
}
