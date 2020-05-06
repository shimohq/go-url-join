package urljoin

import (
	"regexp"
	"strings"
)

// Config is config for joining URL parts
type Config struct {
	// Adds or removes leading slash. Turned on by default.
	LeadingSlash bool
	// Preserves what the leading slash only if it's present on the input.
	KeepLeadingSlash bool
	// Adds or removes trailing slash.
	TrailingSlash bool
	// Preserves what the trailing slash only if it's present on the input.
	KeepTrailingSlash bool
}

var defaultUrlRegExp = regexp.MustCompile(`^(\w+://[^/?]+)?(.*?)(\?.+)?$`)

func normalizeParts(input []string) []string {
	var output []string
	for _, v := range input {
		s := strings.TrimSpace(v)
		if s != "" {
			output = append(output, s)
		}
	}

	return output
}

// Join joins all given URL segments together with default config.
func Join(parts ...string) string {
	return JoinWithConfig(Config{LeadingSlash: true}, parts...)
}

// Join joins all given URL segments together with given config.
func JoinWithConfig(config Config, parts ...string) string {
	parts = normalizeParts(parts)

	var prefix string
	var pathname string
	var suffix string

	joinedParts := strings.Join(parts, "/")

	matched := defaultUrlRegExp.FindStringSubmatch(joinedParts)
	if len(matched) > 0 {
		prefix = matched[1]
		pathname = matched[2]
		suffix = matched[3]
	}

	var hasLeading bool
	var hasTrailing bool
	if suffix != "" {
		hasLeading = regexp.MustCompile(`^//+`).MatchString(pathname)
		hasTrailing = regexp.MustCompile(`//+$`).MatchString(pathname)
	} else {
		hasLeading = regexp.MustCompile(`^/+`).MatchString(pathname)
		hasTrailing = regexp.MustCompile(`/+$`).MatchString(pathname)
	}

	addLeading := config.LeadingSlash || (hasLeading && config.KeepLeadingSlash)
	addTrailing := config.TrailingSlash || (hasTrailing && config.KeepTrailingSlash)

	var u string

	var pathnameParts []string
	for _, part := range strings.Split(pathname, "/") {
		if part != "" {
			pathnameParts = append(pathnameParts, part)
		}
	}

	if len(pathnameParts) > 0 {
		u = strings.Join(pathnameParts, "/")

		if prefix == "" {
			if !strings.HasPrefix(u, "/") {
				if strings.HasPrefix(joinedParts, "//") {
					u = "//" + u
				} else if addLeading {
					u = "/" + u
				}
			}
		} else {
			if !strings.HasPrefix(u, "/") {
				u = prefix + "/" + u
			} else {
				u = prefix + u
			}
		}
	}

	if addTrailing {
		u += "/"
	}

	if u == "" && addLeading {
		u += "/"
	}

	if suffix != "" {
		u += suffix
	}

	return u
}
