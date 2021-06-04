# Change log
This document describes changes between versions of rbxmk. The `imperative`
branch is the latest unreleased version.

## imperative
**Incompatibilities**
- Move rbxmk.newAttrConfig to AttrConfig.new.
- Move rbxmk.newCookie to Cookie.new.
- Move rbxmk.cookiesFrom to Cookie.from.
- Split rbxmk.newDesc into individual constructors.
	- RootDesc.new.
	- ClassDesc.new.
	- PropertyDesc.new.
	- FunctionDesc.new.
	- EventDesc.new.
	- CallbackDesc.new.
	- ParameterDesc.new.
	- TypeDesc.new.
	- EnumDesc.new.
	- EnumItemDesc.new.
- Move os library functions to path library.
	- Move os.clean to path.clean
	- Move os.expand to path.expand
	- Move os.join to path.join
	- Move os.split to path.split
	- os.getenv is unchanged.

**Highlights:**
- Add [path.clean](https://github.com/Anaminus/rbxmk/blob/imperative/doc/libraries.md#user-content-pathclean) function to clean file paths.
- Improve handling of values in the [rbx formats](https://github.com/Anaminus/rbxmk/blob/imperative/doc/formats.md#user-content-rbxl).
- Add "Desc" option to the [rbx formats](https://github.com/Anaminus/rbxmk/blob/imperative/doc/formats.md#user-content-rbxl) to set a descriptor used when encoding or decoding.
- Add "DescMode" option to the [rbx formats](https://github.com/Anaminus/rbxmk/blob/imperative/doc/formats.md#user-content-rbxl) to set how deviations from the descriptor are handled.
- Including `--desc-*` flags causes the "Enum" global variable to be set to the
  enums generated from the resulting descriptor.
- Implement members for the DescAction type.
- Add DescAction.new constructor.
- Add Enum field to rbxmk library to contain enums defined by rbxmk.

**Fixes:**
- Fix missing properties from sym.Properties when instance has a descriptor.
- Fix unexpected conversion failures when getting certain property types while
  an instance has descriptor.
- Fix type conversion error when encoding `desc-patch.json` format.
- Fix type conversion error when encoding Objects with [rbx formats](https://github.com/Anaminus/rbxmk/blob/imperative/doc/formats.md#user-content-rbxl).
- Fix encoding of SharedStrings in rbxlx/rbxmx formats.
- Fix incorrect use of CDATA sections in rbxlx/rbxmx formats.
- Fix interactive command not having descriptor flags.
- Fix Cookie.from returning incorrect values when no cookies are found.

See a [comparison with the previous version][cmp-imperative] for a thorough list of changes.

The [Documentation page][doc-imperative] provides a complete reference for this version of rbxmk.

[doc-imperative]: https://github.com/Anaminus/rbxmk/blob/imperative/doc/README.md#user-content-rbxmk-reference
[cmp-imperative]: https://github.com/Anaminus/rbxmk/compare/v0.7.2...imperative

## v0.7.2
**Fixes:**
- Fix fragments failing to resolve due to case-sensitivity.

See a [comparison with the previous version][cmp-v0.7.2] for a thorough list of changes.

The [Documentation page][doc-v0.7.2] provides a complete reference for this version of rbxmk.

[doc-v0.7.2]: https://github.com/Anaminus/rbxmk/blob/v0.7.2/doc/README.md#user-content-rbxmk-reference
[cmp-v0.7.2]: https://github.com/Anaminus/rbxmk/compare/v0.7.1...v0.7.2

## v0.7.1
**Highlights:**
- Add [Instance\[sym.Properties\]](https://github.com/Anaminus/rbxmk/blob/v0.7.1/doc/types.md#user-content-instancesymproperties) to get or set all properties of an instance at once.
- Add [`rbxmk i`](https://github.com/Anaminus/rbxmk/blob/v0.7.1/doc/README.md#user-content-interactive-command) command to enter an interactive prompt.
- Add flags to [`run`](https://github.com/Anaminus/rbxmk/blob/v0.7.1/doc/README.md#user-content-run-command) command that preset a global descriptor.
	- The `--desc-latest` flag uses the latest API dump from Roblox.
	- The `--desc-file` flag uses a file in the descriptor format (API dump).
	- The `--desc-patch` flag uses a file in the descriptor patch format.
- Add `--libraries` flag to [`run`](https://github.com/Anaminus/rbxmk/blob/v0.7.1/doc/README.md#user-content-run-command) and [`i`](https://github.com/Anaminus/rbxmk/blob/v0.7.1/doc/README.md#user-content-interactive-command) to include and exclude libraries.

**Fixes:**
- Fix [`download-asset`](https://github.com/Anaminus/rbxmk/blob/v0.7.1/doc/README.md#user-content-download-asset-command) flags not working.
- Fix error handling in [`download-asset`](https://github.com/Anaminus/rbxmk/blob/v0.7.1/doc/README.md#user-content-download-asset-command).
- Fix error when getting certain string exprim types.
- Fix path security error when running a file outside of working directory.
- Fix runtime error produced by [os.split](https://github.com/Anaminus/rbxmk/blob/v0.7.1/doc/libraries.md#user-content-ossplit).

See a [comparison with the previous version][cmp-v0.7.1] for a thorough list of changes.

The [Documentation page][doc-v0.7.1] provides a complete reference for this version of rbxmk.

[doc-v0.7.1]: https://github.com/Anaminus/rbxmk/blob/v0.7.1/doc/README.md#user-content-rbxmk-reference
[cmp-v0.7.1]: https://github.com/Anaminus/rbxmk/compare/v0.6.2...v0.7.1

## v0.7.0
Botched release. See v0.7.1.

## v0.6.2
**Highlights:**
- Add [`rbxmk download-asset`](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/README.md#user-content-download-asset-command) command to quickly download an asset from the Roblox website.
- Add [`rbxmk upload-asset`](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/README.md#user-content-upload-asset-command) command to quickly upload an asset to the Roblox website.
- Add [`rbxmk dump`](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/README.md#user-content-dump-command) command to dump the rbxmk Lua API in various formats.
	- Supports generic JSON and minified JSON format.
	- Supports [selene](https://kampfkarren.github.io/selene/) TOML format.
- Add [Instance.Descend](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/types.md#user-content-instancedescend) as an alternative to child indexing, which rbxmk deliberately does not implement.
- Add `--include-root` flag to [run command](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/README.md#user-content-run-command) to include paths as root directories.
- Add `--allow-insecure-paths` flag to [run command](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/README.md#user-content-run-command) to disable path restrictions.
- Add [CFrame.lookAt](https://developer.roblox.com/en-us/api-reference/datatype/CFrame) constructor.
- Implementations of [Axes.new](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/types.md#user-content-axesnew) and [Faces.new](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/types.md#user-content-facesnew) match Roblox API.
	- Previous implementations exist as [Axes.fromComponents](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/types.md#user-content-axesfromcomponents) and [Faces.fromComponents](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/types.md#user-content-facesfromcomponents).
- Implement face fields on [Axes](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/types.md#user-content-axes), matching Roblox API.
- Rename AttrConfig.new to [rbxmk.newAttrConfig](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/libraries.md#user-content-rbxmknewattrconfig).
- Improve speed of [table.clear](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/libraries.md#user-content-tableclear).
- Entries returned by [fs.dir](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/libraries.md#user-content-fsdir) contain only Name and IsDir fields.
	- For large directories, getting files is much faster.
	- Use [fs.stat](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/libraries.md#user-content-fsstat) to get full metadata of a file.

**Fixes:**
- Fix version displayed by rbxmk.
- Fix error when assigning a property to a [DataModel](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/types.md#user-content-datamodel).
- Fix type of [BrickColor](https://developer.roblox.com/en-us/api-reference/datatype/BrickColor) properties decoded by [Roblox XML formats](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/formats.md#user-content-rbxlx).
- Fixes to encoding of [Roblox file formats](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/formats.md#user-content-rbxl).
- Fix [RBXAssetOptions.Cookies](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/types.md#user-content-rbxassetoptions) not being optional.
- Fix missing properties encoded by [Roblox binary formats](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/formats.md#user-content-rbxl) in certain cases.
- Fix [Instance.FindFirstAncestor](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/types.md#user-content-instancefindfirstancestor) behaving as FindFirstAncestorOfClass.
- Fix equality of [Enums](https://developer.roblox.com/en-us/api-reference/datatype/Enums), [Enum](https://developer.roblox.com/en-us/api-reference/datatype/Enum), and [EnumItem](https://developer.roblox.com/en-us/api-reference/datatype/EnumItem) types.
- Fix tostring of [Enums](https://developer.roblox.com/en-us/api-reference/datatype/Enums), [Enum](https://developer.roblox.com/en-us/api-reference/datatype/Enum), and [EnumItem](https://developer.roblox.com/en-us/api-reference/datatype/EnumItem) types.
- Fix handling of nil [Instance](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/types.md#user-content-instance) properties.
- Fix [Instance](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/types.md#user-content-instance) properties not checking inherited classes.
- Fix handling of nil [PhysicalProperties](https://developer.roblox.com/en-us/api-reference/datatype/PhysicalProperties) properties.
- Fix handling of arguments in [fs.mkdir](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/libraries.md#user-content-fsmkdir), [fs.remove](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/libraries.md#user-content-fsremove), and [fs.rename](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/libraries.md#user-content-fsrename).
- Fix [FormatSelectors](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/types.md#user-content-formatselector) being received incorrectly in [clipboard.read](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/libraries.md#user-content-clipboardread) and [clipboard.write](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/libraries.md#user-content-clipboardwrite).
- Fix userdata caching. Immutable types like Vector3 which were equal would incorrectly produce the same userdata. Makes creation of such types faster.
- Fix [os.getenv](https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/libraries.md#user-content-osgetenv) not returning all variables when no value is passed.

**Internal:**
- Automated tests run on Windows in addition to Linux.
- Add tool for automatically incrementing version number.
- Improve documentation.
- Remove concept of "sources"; they're just libraries.

See a [comparison with the previous version][cmp-v0.6.2] for a thorough list of changes.

The [Documentation page][doc-v0.6.2] provides a complete reference for this version of rbxmk.

[doc-v0.6.2]: https://github.com/Anaminus/rbxmk/blob/v0.6.2/doc/README.md#user-content-rbxmk-reference
[cmp-v0.6.2]: https://github.com/Anaminus/rbxmk/compare/v0.5.1...v0.6.2

## v0.6.0, v0.6.1
Botched releases. See v0.6.2.

## v0.5.1
**Internal:**
- Fix automated releases.

See a [comparison with the previous version][cmp-v0.5.1] for a thorough list of changes.

The [Documentation page][doc-v0.5.1] provides a complete reference for this version of rbxmk.

[doc-v0.5.1]: https://github.com/Anaminus/rbxmk/blob/v0.5.1/rbxmk/doc/DOCUMENTATION.md#documentation
[cmp-v0.5.1]: https://github.com/Anaminus/rbxmk/compare/v0.5.0...v0.5.1

## v0.5.0
**Highlights:**
- Improve handling of [HTTP cookies](https://github.com/Anaminus/rbxmk/blob/v0.5.0/doc/types.md#user-content-cookie).
	- Add [rbxmk.cookiesFrom](https://github.com/Anaminus/rbxmk/blob/v0.5.0/doc/libraries.md#user-content-rbxmkcookiesfrom) for getting cookies from known locations.
	- Add [rbxmk.newCookie](https://github.com/Anaminus/rbxmk/blob/v0.5.0/doc/libraries.md#user-content-rbxmknewcookie) for creating a cookie from scratch.
- DataModel metadata is now accessed through the Metadata symbol.
- Rename ["file" source](https://github.com/Anaminus/rbxmk/blob/v0.5.0/doc/sources.md#user-content-fs) to "fs".
- Move [os.dir](https://github.com/Anaminus/rbxmk/blob/v0.5.0/doc/sources.md#user-content-fsdir) and [os.stat](https://github.com/Anaminus/rbxmk/blob/v0.5.0/doc/sources.md#user-content-fsstat) to the [fs library](https://github.com/Anaminus/rbxmk/blob/v0.5.0/doc/sources.md#user-content-fs).
- [fs.dir](https://github.com/Anaminus/rbxmk/blob/v0.5.0/doc/sources.md#user-content-fsdir) and [fs.stat](https://github.com/Anaminus/rbxmk/blob/v0.5.0/doc/sources.md#user-content-fsstat) return nil if the file does not exist.
- Additions to [fs library](https://github.com/Anaminus/rbxmk/blob/v0.5.0/doc/sources.md#user-content-fs).
	- [fs.mkdir](https://github.com/Anaminus/rbxmk/blob/v0.5.0/doc/sources.md#user-content-fsmkdir) for creating directories.
	- [fs.remove](https://github.com/Anaminus/rbxmk/blob/v0.5.0/doc/sources.md#user-content-fsremove) for removing files.
	- [fs.rename](https://github.com/Anaminus/rbxmk/blob/v0.5.0/doc/sources.md#user-content-fsrename) for moving files.
- To reduce the impact of malicious scripts, files can only be accessed by rbxmk from certain locations.
	- The working directory.
	- The directory of the first running script.
	- A temporary directory, accessible via `os.expand("$tmp")`
- Changes to [os.expand](https://github.com/Anaminus/rbxmk/blob/v0.5.0/doc/libraries.md#user-content-osexpand).
	- Add the `$root_script_directory` variable, which expands to the directory of the first running script.
	- The `$temp_directory` variable now points to a temporary directory that is unique per run of rbxmk.

**Fixes:**
- Fix garbled error messages.

**Internal:**
- Implement automated releases.

See a [comparison with the previous version][cmp-v0.5.0] for a thorough list of changes.

The [Documentation page][doc-v0.5.0] provides a complete reference for this version of rbxmk.

[doc-v0.5.0]: https://github.com/Anaminus/rbxmk/blob/v0.5.0/doc/README.md#user-content-rbxmk-reference
[cmp-v0.5.0]: https://github.com/Anaminus/rbxmk/compare/v0.4.0...v0.5.0

## v0.4.0
**Highlights:**
- The [command-line API](https://github.com/Anaminus/rbxmk/tree/v0.4.0/doc#user-content-command-line) now uses sub-commands.
	- `rbxmk run` runs scripts as before.
	- `rbxmk help` displays help.
	- `rbxmk version` displays the current version.
- Implement [Instance attributes](https://github.com/Anaminus/rbxmk/tree/v0.4.0/doc#user-content-attributes).
	- [Instance.GetAttribute](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/types.md#user-content-instancegetattribute)
	- [Instance.GetAttributes](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/types.md#user-content-instancegetattributes)
	- [Instance.SetAttribute](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/types.md#user-content-instancesetattribute).
	- Additional [SetAttributes](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/types.md#user-content-instancesetattributes) method for efficiently setting all attributes at once.
- Add inheritance-based methods to [Instance](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/types.md#user-content-instance).
	- [Instance.IsA](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/types.md#user-content-instanceisa)
	- [Instance.FindFirstChildWhichIsA](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/types.md#user-content-instancefindfirstchildwhichisa)
	- [Instance.FindFirstAncestorWhichIsA](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/types.md#user-content-instancefindfirstancestorwhichisa)
- [Format configuration](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/types.md#user-content-formatselector): in APIs where a format string could be passed, a table can be passed instead, which can contain additional options that configure the format.
- The [http source](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/sources.md#user-content-http) has been rewritten.
	- Has single [http.request](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/sources.md#user-content-httprequest) function.
	- Handles CSRF tokens automatically.
- Add [rbxassetid source](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/sources.md#user-content-rbxassetid) for accessing assets on the Roblox website.
- Add [clipboard source](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/sources.md#user-content-clipboard) for accessing the operating system's clipboard.
	- **Currently available only on Windows.**
- Add [os.stat](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/libraries.md#user-content-osstat) function for getting the metadata of a file.
- Additions to the [math library](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/libraries.md#user-content-math).
	- [math.log](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/libraries.md#user-content-mathlog) from Lua 5.2.
	- [math.clamp](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/libraries.md#user-content-mathclamp) from Roblox.
	- [math.round](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/libraries.md#user-content-mathround) from Roblox.
	- [math.sign](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/libraries.md#user-content-mathsign) from Roblox.
- Additions to the [table library](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/libraries.md#user-content-table).
	- [table.pack](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/libraries.md#user-content-tablepack) from Lua 5.2.
	- [table.unpack](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/libraries.md#user-content-tableunpack) from Lua 5.2.
	- [table.move](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/libraries.md#user-content-tablemove) from Lua 5.3.
	- [table.clear](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/libraries.md#user-content-tableclear) from Roblox.
	- [table.create](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/libraries.md#user-content-tablecreate) from Roblox.
	- [table.find](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/libraries.md#user-content-tablefind) from Roblox.
- Additions to the [string library](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/libraries.md#user-content-string).
	- [string.split](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/libraries.md#user-content-stringsplit) from Roblox.
- Add [rbxmk.formatCanDecode](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/libraries.md#user-content-rbxmkformatcandecode) function for getting whether a given format can decode into a given type.
- Add [json format](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/formats.md#user-content-json) for decoding JSON data generically, similar to HttpService.JSONEncode/JSONDecode.
- Add [csv format](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/formats.md#user-content-csv) for converting general CSV data.
- Add [l10n.csv format](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/formats.md#user-content-l10ncsv) for converting localization data.
- [Instance.new](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/types.md#user-content-instancenew) is allowed to create classes with the NotCreatable tag.
- Document `_RBXMK_VERSION` global variable.
- Remove rbxmk.readSource and rbxmk.writeSource functions.

**Fixes:**
- Fix sorting of members returned by [ClassDesc.Members](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/types.md#user-content-classdescmembers).
- Fix error with [EnumDesc.AddItem](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/types.md#user-content-enumdescadditem).
- Fix error with [RootDesc.AddEnum](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/types.md#user-content-rootdescaddenum).
- Fix error when indexing Enums or calling GetEnums.
- Fix Enums.GetEnums returning no values.
- Fix issues with cloned [Instances](https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/types.md#user-content-instance).
- Fix string and double types being returned as empty userdata.
- Fix strings not having a metatable.

See a [comparison with the previous version][cmp-v0.4.0] for a thorough list of changes.

The [Documentation page][doc-v0.4.0] provides a complete reference for this version of rbxmk.

[doc-v0.4.0]: https://github.com/Anaminus/rbxmk/blob/v0.4.0/doc/README.md#user-content-rbxmk-reference
[cmp-v0.4.0]: https://github.com/Anaminus/rbxmk/compare/v0.3.0...v0.4.0

## v0.3.0
**Imperative Mode** is a complete rewrite of rbxmk from the ground up.
- Versus the previous quasi-declarative Lua API, the new API is completely imperative.
- The API is designed to emulate the Roblox Lua API.

The [Documentation page][doc-v0.3.0] provides a complete reference for this version of rbxmk.

[doc-v0.3.0]: https://github.com/Anaminus/rbxmk/blob/v0.3.0/rbxmk/doc/DOCUMENTATION.md#documentation
