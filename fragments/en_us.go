//go:build lang_en_us

package fragments

import "embed"

//go:embed en-us
var en_us embed.FS

func init() { register("en-us", en_us) }
