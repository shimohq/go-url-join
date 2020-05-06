package urljoin

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJoin(t *testing.T) {
	a := assert.New(t)

	a.Equal(Join("foo", "bar"), "/foo/bar")
	a.Equal(Join("/foo/", "/bar/"), "/foo/bar")
	a.Equal(Join("foo", "", "bar"), "/foo/bar")
	a.Equal(Join("https://wikipedia.org", "foo"), "https://wikipedia.org/foo")
	a.Equal(Join("https://wikipedia.org/", "foo"), "https://wikipedia.org/foo")
	a.Equal(Join("foo", "bar?a=b"), "/foo/bar?a=b")
	a.Equal(Join("//wikipedia.org", "foo"), "//wikipedia.org/foo")
	a.Equal(Join("wikipedia.org", "foo"), "/wikipedia.org/foo")
	a.Equal(Join(""), "/")
}

func TestJoinWithConfig (t *testing.T) {
	a := assert.New(t)

	a.Equal(JoinWithConfig(Config{LeadingSlash: true}, "foo", "bar"), "/foo/bar")
	a.Equal(JoinWithConfig(Config{LeadingSlash: false}, "foo", "bar"), "foo/bar")
	a.Equal(JoinWithConfig(Config{LeadingSlash: true}, "/foo/", "/bar/"), "/foo/bar")
	a.Equal(JoinWithConfig(Config{LeadingSlash: false}, "/foo/", "/bar/"), "foo/bar")
	a.Equal(JoinWithConfig(Config{LeadingSlash: false}, "foo", "", "bar"), "foo/bar")
	a.Equal(JoinWithConfig(Config{LeadingSlash: false}, "wikipedia.org", "bar"), "wikipedia.org/bar")
	a.Equal(JoinWithConfig(Config{LeadingSlash: false}, ""), "")

	a.Equal(JoinWithConfig(
		Config{LeadingSlash: false, KeepLeadingSlash: true},
		"/foo/", "/bar/"),
		"/foo/bar",
	)

	a.Equal(
		JoinWithConfig(Config{LeadingSlash: true},
		"https://wikipedia.org", "foo"),
		"https://wikipedia.org/foo",
	)

	a.Equal(JoinWithConfig(Config{LeadingSlash: true}, "foo", "bar?a=b"), "/foo/bar?a=b")
	a.Equal(JoinWithConfig(Config{LeadingSlash: false}, "foo", "bar?a=b"), "foo/bar?a=b")

	a.Equal(JoinWithConfig(Config{TrailingSlash: true}, "foo", "bar/"), "foo/bar/")
	a.Equal(JoinWithConfig(Config{TrailingSlash: false}, "foo", "bar"), "foo/bar")
	a.Equal(JoinWithConfig(Config{TrailingSlash: true}, "foo", "bar"), "foo/bar/")
	a.Equal(JoinWithConfig(Config{TrailingSlash: true}, "foo", "bar?a=b"), "foo/bar/?a=b")

	a.Equal(JoinWithConfig(
		Config{TrailingSlash: false, KeepTrailingSlash: true},
		"foo", "bar/"),
		"foo/bar/",
	)

	a.Equal(
		JoinWithConfig(Config{},
		"//wikipedia.org", "foo"),
		"//wikipedia.org/foo",
	)
	a.Equal(
		JoinWithConfig(Config{},
			"https://wikipedia.org", "foo"),
		"https://wikipedia.org/foo",
	)

	a.Equal(JoinWithConfig(Config{LeadingSlash: true}, "https://wikipedia.org", "foo"), "https://wikipedia.org/foo")
	a.Equal(JoinWithConfig(Config{LeadingSlash: false}, "https://wikipedia.org", "foo"), "https://wikipedia.org/foo")
	a.Equal(JoinWithConfig(Config{LeadingSlash: true}, "//wikipedia.org", "foo"), "//wikipedia.org/foo")
	a.Equal(JoinWithConfig(Config{LeadingSlash: false}, "//wikipedia.org", "foo"), "//wikipedia.org/foo")
}
