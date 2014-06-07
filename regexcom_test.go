package regexcom

import (
	"log"
	"testing"
)

type patternTest struct {
	name    string   // Name pattern to test
	typ     string   // Type of test
	pattern *Pattern // Actual patern object
	value   string   // The value to run the regex on
	doMatch bool     // Wheather or not the regex should match
}

var (
	tests = []patternTest{
		// Numeric
		{"Numeric", "Simple", Numeric, "1", true},
		{"Numeric", "Hard", Numeric, "1947189", true},
		{"Numeric", "FalseSimple", Numeric, "a", false},
		{"Numeric", "FalseHard", Numeric, "ABHJ4A642", false},

		// Alpha
		{"Alpha", "Simple", Alpha, "a", true},
		{"Alpha", "Hard", Alpha, "afAfeAs", true},
		{"Alpha", "FalseSimple", Alpha, "1", false},
		{"Alpha", "FalseHard", Alpha, "aF4!f4Am,", false},

		// AlphaNumeric
		{"AlphaNumeric", "Simple", AlphaNumeric, "1", true},
		{"AlphaNumeric", "Hard", AlphaNumeric, "1947189", true},
		{"AlphaNumeric", "FalseSimple", AlphaNumeric, "a1%", false},
		{"AlphaNumeric", "FalseHard", AlphaNumeric, "AB^3HJ4A642", false},

		// AlphaCapsOnly
		{"AlphaCapsOnly", "Simple", AlphaCapsOnly, "A", true},
		{"AlphaCapsOnly", "Hard", AlphaCapsOnly, "ASDJHGASD", true},
		{"AlphaCapsOnly", "FalseSimple", AlphaCapsOnly, "a", false},
		{"AlphaCapsOnly", "FalseHard", AlphaCapsOnly, "ADsa&", false},

		// AlphaNumericCapsOnly
		{"AlphaNumericCapsOnly", "Simple", AlphaNumericCapsOnly, "A1", true},
		{"AlphaNumericCapsOnly", "Hard", AlphaNumericCapsOnly, "ASD5435JHGASD", true},
		{"AlphaNumericCapsOnly", "FalseSimple", AlphaNumericCapsOnly, "a1", false},
		{"AlphaNumericCapsOnly", "FalseHard", AlphaNumericCapsOnly, "a534&", false},

		// Url
		{"Url", "Simple", Url, "b.ca", true},
		{"Url", "Hard", Url, "example.com", true},
		{"Url", "HTTP", Url, "http://example.com", true},
		{"Url", "HTTPS", Url, "https://example.com", true},
		{"Url", "FTP", Url, "ftp://example.com", true},
		{"Url", "WWW-HTTP", Url, "http://wwww.example.com", true},
		{"Url", "WWW-HTTPS", Url, "https://wwww.example.com", true},
		{"Url", "WWW-FTP", Url, "ftp://wwww.example.com", true},
		{"Url", "FalseSimple", Url, ".com", false},
		{"Url", "FalseHard", Url, "://example.com", false},
		{"Url", "FullMissingTLD", Url, "http://example", false},

		// Email
		{"Email", "Simple", Email, "a@b.cd", true},
		{"Email", "Hard", Email, "user@example.com", true},
		{"Email", "FalseSimple", Email, "@b.cd", false},
		{"Email", "FalseHard", Email, "user@example", false},

		// HashtagHex
		{"HashtagHex", "Simple", HashtagHex, "#afe", true},
		{"HashtagHex", "Hard", HashtagHex, "#a43f1c", true},
		{"HashtagHex", "FalseSimple", HashtagHex, "afe", false},
		{"HashtagHex", "FalseHard", HashtagHex, "AEF14D", false},

		// ZeroXHex
		{"ZeroXHex", "Simple", ZeroXHex, "0xa", true},
		{"ZeroXHex", "Hard", ZeroXHex, "0xF14D", true},
		{"ZeroXHex", "FalseSimple", ZeroXHex, "0x", false},
		{"ZeroXHex", "FalseHard", ZeroXHex, "a41fed", false},
		{"ZeroXHex", "FalseUpperLower", ZeroXHex, "0xaDef1", false},

		// IPv4
		{"IPv4", "Simple", IPv4, "0.0.0.0", true},
		{"IPv4", "Hard", IPv4, "134.10.255.0", true},
		{"IPv4", "FalseSimple", IPv4, "1.1", false},
		{"IPv4", "FalseOver255", IPv4, "256.123.3.450", false},

		// IPv6
		{"IPv6", "Simple", IPv6, "0:0:0:0:0:0:0:0", true},
		{"IPv6", "Hard", IPv6, "2001:0db8:0000:0000:1000:ff00:0042:832a", true},
		{"IPv6", "FalseSimple", IPv6, "0:", false},
		{"IPv6", "FalseOver255", IPv6, "A:A:A:A:A:A::", false},
	}
)

func runAllTests(t *testing.T) {
	for _, test := range tests {
		if test.pattern.Match(test.value) != test.doMatch {
			log.Printf("\tFAIL: Test_%v%v", test.name, test.typ)
			if test.doMatch == true {
				t.Errorf("[%v%v] Failed to match '%v' as %v", test.name, test.typ, test.value, test.name)
			} else {
				t.Errorf("[%v%v] Accidentally matched '%v' as %v", test.name, test.typ, test.value, test.name)
			}
		}
	}
}

func Test_RunAll(t *testing.T) {
	runAllTests(t)
}
