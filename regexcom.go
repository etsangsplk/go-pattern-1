package regexcom

import (
	"regexp"
)

var (
	Numeric              = NewPattern(`^(\d+)$`)
	AlphaNumeric         = NewPattern("^([0-9A-Za-z]+)$")
	Alpha                = NewPattern("^([A-Za-z]+)$")
	AlphaCapsOnly        = NewPattern("^([A-Z]+)$")
	AlphaNumericCapsOnly = NewPattern("^([0-9A-Z]+)$")
	Url                  = NewPattern(`^((https?|ftps?):\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`)
	Email                = NewPattern(`^([a-z0-9_\.-]+)@([\da-z\.-]+)\.([a-z\.]{2,6})$`)
	HashtagHex           = NewPattern(`^#([a-f0-9]{6}|[a-f0-9]{3})$`)
	ZeroXHex             = NewPattern(`^0x([a-f0-9]+|[A-F0-9]+)$`)
	IPv4                 = NewPattern(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)
	IPv6                 = NewPattern(`^([0-9A-Fa-f]{0,4}:){2,7}([0-9A-Fa-f]{1,4}$|((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})$`)
)

// Represents the General pattern type
type Pattern struct {
	pattern string
	reg     *regexp.Regexp
}

func NewPattern(pattern string) *Pattern {
	p := &Pattern{pattern: pattern}
	p.reg = regexp.MustCompile(p.pattern)
	return p
}

// Returns true if s matches the compiled regex pattern
func (p *Pattern) Match(s string) bool {
	return p.reg.MatchString(s)
}
