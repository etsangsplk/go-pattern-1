package pattern

import (
	"regexp"
)

var (
	Numeric              = New(`^(\d+)$`)
	AlphaNumeric         = New("^([0-9A-Za-z]+)$")
	Alpha                = New("^([A-Za-z]+)$")
	AlphaCapsOnly        = New("^([A-Z]+)$")
	AlphaNumericCapsOnly = New("^([0-9A-Z]+)$")
	Url                  = New(`^((http?|https?|ftps?):\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`)
	Email                = New(`^(.+@([\da-z\.-]+)\.([a-z\.]{2,6}))$`)
	HashtagHex           = New(`^#([a-f0-9]{6}|[a-f0-9]{3})$`)
	ZeroXHex             = New(`^0x([a-f0-9]+|[A-F0-9]+)$`)
	IPv4                 = New(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)
	IPv6                 = New(`^([0-9A-Fa-f]{0,4}:){2,7}([0-9A-Fa-f]{1,4}$|((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})$`)
)

// Represents the General pattern type
type Pattern struct {
	pat string
	reg *regexp.Regexp
}

func New(pat string) *Pattern {
	p := &Pattern{pat: pat}
	p.reg = regexp.MustCompile(p.pat)
	return p
}

// Returns true if s matches the compiled regex pattern
func (p *Pattern) Match(s string) bool {
	return p.reg.MatchString(s)
}
