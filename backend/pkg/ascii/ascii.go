package ascii

import "unicode/utf8"

// IsASCII checks whether all characters are ASCII which is equivalent to [\x20-\x7E].
func IsASCII(str string) bool {
	// reference: https://qiita.com/catatsuy/items/7a9773f9ea3db7069fc1
	return utf8.ValidString(str) && utf8.RuneCountInString(str) == len(str)
}
