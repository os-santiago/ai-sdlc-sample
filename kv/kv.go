// Package kv parses simple "key=value" lists.
package kv

import (
	"errors"
	"strings"
)

// ErrSyntax reports a malformed pair.
var ErrSyntax = errors.New("kv: malformed pair")

// Parse reads pairs separated by ";" such as "a=1;b=2" into a map.
func Parse(s string) (map[string]string, error) {
	out := map[string]string{}
	for _, pair := range strings.Split(s, ";") {
		k, v, ok := strings.Cut(pair, "=")
		if !ok {
			return nil, ErrSyntax
		}
		out[k] = v
	}
	return out, nil
}
