package pfx

import (
	"fmt"
	"runtime"
	"strings"
)

var FullyQualifiedPath = false

// Err consumes an error, a string, or nil, and produces an error message prefixed
// with the name of the function that called it (or nil).
func Err(err interface{}) error {
	switch o := err.(type) {
	case string:
		return e(fmt.Errorf("%s", o))
	case error:
		return e(o)
	default:
		return nil
	}
}

// e returns an error, prefixed with the name of the function that
// triggered it. Originally by StackOverflow user svenwltr:
// http://stackoverflow.com/a/38551362/199475
func e(err error) error {
	pc, _, _, _ := runtime.Caller(2)

	fr := runtime.CallersFrames([]uintptr{pc})
	namer, _ := fr.Next()
	name := namer.Function

	if !FullyQualifiedPath {
		fn := strings.Split(name, "/")
		if len(fn) > 0 {
			return fmt.Errorf("%s: %w", fn[len(fn)-1], err)
		}
	}

	return fmt.Errorf("%s: %w", name, err)
}
