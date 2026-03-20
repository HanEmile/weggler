package runner

import (
	"regexp"
	"strings"
)

// weggli SGR palette — these are the codes it actually emits:
//   1  = bold (matched variable names / keywords)
//   31 = red      32 = green    33 = yellow
//   34 = blue     35 = magenta  36 = cyan
//   0  = reset
//
// We convert them to inline <span style="…"> tags so the browser terminal
// renders them properly without raw escape characters.

var sgrRe = regexp.MustCompile(`\x1b\[([0-9;]*)[mKHFABCDJsu]`)

// ansiToHTML converts ANSI SGR escape sequences to HTML <span> tags.
// Non-SGR cursor/erase sequences are dropped entirely.
// The result is safe to inject into innerHTML inside a <pre>/<div>.
func ansiToHTML(s string) string {
	var b strings.Builder
	b.Grow(len(s) + 128)

	// Stack of open spans so we can close them properly on reset.
	open := 0

	closeAll := func() {
		for open > 0 {
			b.WriteString("</span>")
			open--
		}
	}

	last := 0
	for _, m := range sgrRe.FindAllStringSubmatchIndex(s, -1) {
		// m[0]:m[1] = full match, m[2]:m[3] = capture group (the numeric part)
		// Write literal text before this escape, HTML-escaped
		b.WriteString(htmlEscape(s[last:m[0]]))
		last = m[1]

		params := s[m[2]:m[3]]
		// Only handle 'm' (SGR) — the regex also matches cursor codes but we
		// simply drop them by not writing anything.
		if s[m[1]-1] != 'm' {
			continue
		}

		if params == "" || params == "0" {
			closeAll()
			continue
		}

		style := sgrToStyle(params)
		if style != "" {
			b.WriteString(`<span style="`)
			b.WriteString(style)
			b.WriteString(`">`)
			open++
		}
		// Unknown codes: silently drop.
	}

	// Remaining literal text
	b.WriteString(htmlEscape(s[last:]))
	closeAll()
	return b.String()
}

// StripANSI removes all ANSI escape sequences, returning plain text.
// Used for finding parsing where we need clean file:line headers.
func StripANSI(s string) string { return sgrRe.ReplaceAllString(s, "") }

func sgrToStyle(params string) string {
	// Handle combined codes like "1;32"
	bold := false
	color := ""
	bgColor := ""
	for _, part := range strings.Split(params, ";") {
		switch part {
		case "1":
			bold = true
		case "2": // dim — skip
		case "3": // italic
		case "4": // underline — skip
		// Foreground colors (standard)
		case "30":
			color = "#4a5568" // dark grey
		case "31":
			color = "#fc8181" // red
		case "32":
			color = "#68d391" // green  ← weggli uses this for matched captures
		case "33":
			color = "#f6e05e" // yellow
		case "34":
			color = "#63b3ed" // blue
		case "35":
			color = "#b794f4" // magenta
		case "36":
			color = "#76e4f7" // cyan
		case "37":
			color = "#e2e8f0" // white
		// Bright foreground
		case "90":
			color = "#718096"
		case "91":
			color = "#fc8181"
		case "92":
			color = "#68d391"
		case "93":
			color = "#faf089"
		case "94":
			color = "#76e4f7"
		case "95":
			color = "#fbb6ce"
		case "96":
			color = "#81e6d9"
		case "97":
			color = "#f7fafc"
		// Background colors (we tone them down to subtle backgrounds)
		case "40":
			bgColor = "rgba(74,85,104,.3)"
		case "41":
			bgColor = "rgba(252,129,129,.15)"
		case "42":
			bgColor = "rgba(104,211,145,.15)"
		case "43":
			bgColor = "rgba(246,224,94,.12)"
		case "44":
			bgColor = "rgba(99,179,237,.15)"
		case "45":
			bgColor = "rgba(183,148,244,.15)"
		case "46":
			bgColor = "rgba(118,228,247,.12)"
		}
	}
	if !bold && color == "" && bgColor == "" {
		return ""
	}
	var sb strings.Builder
	if bold {
		sb.WriteString("font-weight:700;")
	}
	if color != "" {
		sb.WriteString("color:")
		sb.WriteString(color)
		sb.WriteString(";")
	}
	if bgColor != "" {
		sb.WriteString("background:")
		sb.WriteString(bgColor)
		sb.WriteString(";border-radius:2px;padding:0 1px;")
	}
	return sb.String()
}

func htmlEscape(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	return s
}
