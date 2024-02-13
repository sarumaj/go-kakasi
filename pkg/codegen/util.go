package codegen

import (
	"bufio"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// decodeEscapes decodes escape sequences in a string.
func decodeEscapes(s string) string {
	// Unescape Unicode sequences
	reUnicode := regexp.MustCompile(`\\[uU][0-9a-fA-F]{4,8}`)
	s = reUnicode.ReplaceAllStringFunc(s, func(match string) string {
		r, err := strconv.ParseInt(match[2:], 16, 32)
		if err != nil {
			return match // Return the original match if there's an error
		}
		return string(rune(r))
	})

	// Unescape 2-digit hex escapes
	reHex := regexp.MustCompile(`\\x[0-9a-fA-F]{2}`)
	s = reHex.ReplaceAllStringFunc(s, func(match string) string {
		r, err := strconv.ParseUint(match[2:], 16, 8)
		if err != nil {
			return match // Return the original match if there's an error
		}

		return string(rune(r))
	})

	// Unescape octal escapes
	reOctal := regexp.MustCompile(`\\[0-7]{1,3}`)
	s = reOctal.ReplaceAllStringFunc(s, func(match string) string {
		r, err := strconv.ParseUint(match[1:], 8, 8)
		if err != nil {
			return match // Return the original match if there's an error
		}

		return string(rune(r))
	})

	// Unescape basic escape sequences
	replacements := []string{
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
	}

	return strings.NewReplacer(replacements...).Replace(s)
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
func mapGet[K comparable, T any](m map[K]T, k K) T {
	return m[k]
}

// mapHas checks if a map has a key.
func mapHas[K comparable, T any](m map[K]T, k K) bool {
	_, ok := m[k]
	return ok
}

// mapKeys returns the keys of a map.
func mapKeys[K comparable, T any](m map[K]T) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// mapSet sets a value in a map.
func mapSet[K comparable, T any](m map[K]T, k K, v T) map[K]T {
	m[k] = v
	return m
}

// traverseFile reads a file line by line and sends each line to a channel.
// Empty lines and lines starting with ";;" are ignored.
// The function returns a channel that will be closed when the file has been fully read.
// It also takes a context that can be used to cancel the operation.
func traverseFile(ctx context.Context, in *os.File) <-chan string {
	sc := bufio.NewScanner(in)
	sc.Split(bufio.ScanLines)
	lines := make(chan string, 1)

	go func(sc *bufio.Scanner, lines chan<- string) {
		for sc.Scan() {
			select {

			case <-ctx.Done():
				close(lines)
				return

			default:
				switch line := strings.TrimSpace(sc.Text()); {

				case len(line) == 0, strings.HasPrefix(line, ";;"):
					continue

				default:
					lines <- decodeEscapes(line)
				}
			}

		}

		close(lines)
	}(sc, lines)

	return lines
}
