package hashkey

import (
	"fmt"
	"hash/fnv"
)

// String returns a hashkey for the given string
func String(s string) (string, error) {
	h := fnv.New64a()
	if _, err := h.Write([]byte(s)); err != nil {
		return "", fmt.Errorf("unable to create new hashkey from \"%s\": %w", s, err)
	}
	return fmt.Sprintf("%x", h.Sum64()), nil
}

func Uint64(s string) (uint64, error) {
	h := fnv.New64a()
	if _, err := h.Write([]byte(s)); err != nil {
		return 0, fmt.Errorf("unable to create new hashkey from \"%s\": %w", s, err)
	}
	return h.Sum64(), nil
}
