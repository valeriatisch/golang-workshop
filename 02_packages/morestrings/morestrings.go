package morestrings

// ReverseRunes returns its argument string reversed rune-wise left to right.
// from: https://go.dev/doc/code
func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Idiomatic Go
// In Go it is considered idiomatic to put a comment before an exported function declaration.
// This comment should be a complete sentence that begins with the name of the function.
