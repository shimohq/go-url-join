# go-url-join

Like `path.Join()` but for a URL. Inspired by [proper-url-join](https://github.com/moxystudio/js-proper-url-join).

## Installation

```
go get github.com/shimohq/go-url-join
```

## Example

```go
package main

import (
  "github.com/shimohq/go-url-join"
)

func main() {
    urljoin.Join("foo", "bar") // "/foo/bar"
    urljoin.Join("//wikipedia.org", "foo") // "//wikipedia.org/foo"
    urljoin.Join("https://wikipedia.org", "foo") // "https://wikipedia.org/foo"

    urljoin.JoinWithConfig(
        urljoin.Config{TrailingSlash: true},
        "foo", "bar",
    ) // "/foo/bar"

    urljoin.JoinWithConfig(
        urljoin.Config{LeadingSlash: false, KeepLeadingSlash: false},
        "/foo", "bar",
    ) // "foo/bar"

    urljoin.JoinWithConfig(
        urljoin.Config{LeadingSlash: false, KeepLeadingSlash: true},
        "/foo", "bar",
    ) // "/foo/bar"
}
```
