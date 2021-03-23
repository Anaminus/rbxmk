// +build !lang_no_default

package fragments

import "embed"

//go:embed en-us
var _default embed.FS

func init() { register("en-us", _default) }
