// The fragments package contains translations of documentation fragments.
//
// To include a translation, the package must be built with a number of
// lang_$LOCALE tag, where $LOCALE is the name of the translation. The name is
// lower-case, with dashes replaced by underscores. For example,
//
//     -tags lang_en_us
//
// An "en-us" translation is included by default. To exclude it, use the
// "lang_no_default" tag.
package fragments

import "embed"

// Languages contains locale names mapped to embedded fragment files.
var Languages = map[string]embed.FS{}

func register(name string, value embed.FS) { Languages[name] = value }
