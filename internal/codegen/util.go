package codegen

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	json "github.com/goccy/go-json"
	ordered "github.com/wk8/go-ordered-map/v2"
)

// Escape sequences understood by decodeEscapes. The patterns and the replacer
// are compiled once: decodeEscapes runs on every line of every source
// dictionary, which is hundreds of thousands of lines.
var (
	reUnicode = regexp.MustCompile(`\\[uU][0-9a-fA-F]{4,8}`)
	reHex     = regexp.MustCompile(`\\x[0-9a-fA-F]{2}`)
	reOctal   = regexp.MustCompile(`\\[0-7]{1,3}`)

	basicEscapes = strings.NewReplacer(
		`\\`, `\`,
		`\"`, `"`,
		`\'`, `'`,
		`\n`, "\n",
		`\r`, "\r",
		`\t`, "\t",
		`\b`, "\b",
		`\f`, "\f",
		`\a`, "\a",
		`\v`, "\v",
	)
)

// decodeEscapes decodes escape sequences in a string.
func decodeEscapes(s string) string {
	// Unescape Unicode sequences
	s = reUnicode.ReplaceAllStringFunc(s, func(match string) string {
		r, err := strconv.ParseInt(match[2:], 16, 32)
		if err != nil {
			return match // Return the original match if there's an error
		}

		return string(rune(r))
	})

	// Unescape 2-digit hex escapes
	s = reHex.ReplaceAllStringFunc(s, func(match string) string {
		r, err := strconv.ParseUint(match[2:], 16, 8)
		if err != nil {
			return match // Return the original match if there's an error
		}

		return string(rune(r))
	})

	// Unescape octal escapes
	s = reOctal.ReplaceAllStringFunc(s, func(match string) string {
		r, err := strconv.ParseUint(match[1:], 8, 8)
		if err != nil {
			return match // Return the original match if there's an error
		}

		return string(rune(r))
	})

	return basicEscapes.Replace(s)
}

// deref dereferences a pointer.
// If the pointer is nil, the zero value of the type is returned.
func deref[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}

	return *v
}

// dump writes a value to a file in JSON format.
// The file will be created if it doesn't exist, and truncated if it does.
// The directory structure will be created if it doesn't exist a priori.
func dumpJSON(dst string, v any, indent string) error {
	_ = os.MkdirAll(filepath.Dir(dst), os.ModePerm)
	o, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}

	defer o.Close()

	enc := json.NewEncoder(o)
	enc.SetIndent("", indent)
	return enc.Encode(v)
}

// mapGet gets a value from a map.
func mapGet[K comparable, T any](m ordered.OrderedMap[K, T], k K) T {
	return m.Value(k)
}

// mapHas checks if a map has a key.
func mapHas[K comparable, T any](m ordered.OrderedMap[K, T], k K) bool {
	_, ok := m.Get(k)
	return ok
}

// mapIter returns an iterator for a map.
func mapIter[K comparable, V any](m ordered.OrderedMap[K, V]) func() (K, V, bool) {
	var current *ordered.Pair[K, V] = m.Oldest()
	return func() (K, V, bool) {
		if current == nil {
			var zeroK K
			var zeroV V
			return zeroK, zeroV, false
		}

		defer func() { current = current.Next() }()
		return current.Key, current.Value, true
	}
}

// mapKeys returns the keys of a map.
func mapKeys[K comparable, T any](m ordered.OrderedMap[K, T]) []K {
	keys := make([]K, 0, m.Len())
	for p := m.Oldest(); p != nil; p = p.Next() {
		keys = append(keys, p.Key)

	}

	return keys
}

// mapLen returns the length of a map.
func mapLen[K comparable, T any](m ordered.OrderedMap[K, T]) int {
	return m.Len()
}

// mapSet sets a value in a map.
func mapSet[K comparable, T any](m *ordered.OrderedMap[K, T], k K, v T) *ordered.OrderedMap[K, T] {
	_, _ = m.Set(k, v)
	return m
}

// traverseFile reads a file line by line and passes each line to fn.
// Empty lines and lines starting with ";;" are ignored, and escape sequences
// are decoded before fn is called. It stops and returns the first error
// reported by fn or by the underlying scanner.
func traverseFile(src string, fn func(line string) error) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}

	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		switch line := strings.TrimSpace(sc.Text()); {

		case len(line) == 0, strings.HasPrefix(line, ";;"):
			continue

		default:
			if err := fn(decodeEscapes(line)); err != nil {
				return err
			}

		}
	}

	return sc.Err()
}
