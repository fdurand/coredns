// +build gofuzz

package cache

import (
	"github.com/fdurand/coredns/plugin/pkg/fuzz"
)

// Fuzz fuzzes cache.
func Fuzz(data []byte) int {
	return fuzz.Do(New(), data)
}
