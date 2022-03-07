# API Dump patch
The [dump.desc-patch.json](dump.desc-patch.json) file patches Roblox's API dump
to include missing items.

The patch can be applied through the command line with the `--desc-patch` flag
following another descriptor flag:

```bash
rbxmk run --desc-latest --desc-patch dump.desc-patch.json script.lua
```

It can also be applied within a script with the Patch method:

```lua
local desc = fs.read("dump.desc.json")
local patch = fs.read("dump.desc-patch.json")
desc:Patch(patch)
```

The [patch.lua](patch.lua) script can be used to overwrite an existing
descriptor file to include a patch.
