package textutil

import "regexp"

var dashesRegexp = regexp.MustCompile(`-{2,}`)

// CollapseDashes replaces every run of two or more - with a single -.
func CollapseDashes(s string) string {
	return dashesRegexp.ReplaceAllString(s, "-")
}
