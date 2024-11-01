package conf

import (
	"strings"

	p "github.com/Oudwins/zog/internals"
)

// IsZeroValueFunc is a function that determines if a value is zero based on the logic of Zog which may be different than Go's default behavior.
type IsZeroValueFunc = func(val any) bool

// IsZeroValue is a map of functions to determine if a value is zero.
// This is used to determine how to handle required, default and catch schema properties.
var DefaultParseIsZeroValue = struct {
	Bool    IsZeroValueFunc
	String  IsZeroValueFunc
	Int     IsZeroValueFunc
	Float64 IsZeroValueFunc
	Time    IsZeroValueFunc
	Slice   IsZeroValueFunc
	// Struct  IsZeroValueFunc
}{
	Bool: func(val any) bool {
		iszero := p.IsZeroValue(val)
		if !iszero {
			return iszero
		}
		// We want to know if its a bool because we want to accept false as a value but not "" or nil
		_, ok := val.(bool)
		return !ok
	},
	String: func(val any) bool {
		s, ok := val.(string)
		if ok {
			// Considers "   " to be a zero value
			return p.IsZeroValue(strings.TrimSpace(s))
		}
		return p.IsZeroValue(val)
	},
	Int:     p.IsZeroValue,
	Float64: p.IsZeroValue,
	Time:    p.IsZeroValue,
	Slice:   p.IsZeroValue,
	//Struct:  p.IsZeroValue,
}

// ParseIsZeroValue is a map of functions to determine if a value is zero.
// This is used to determine how to handle required, default and catch schema properties.
var ParseIsZeroValue = DefaultParseIsZeroValue
