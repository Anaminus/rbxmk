# Summary
Execute a script.

# Arguments
	[ FLAGS ] FILE [ VALUE... ]

# Description
The **run** command receives a file to be executed as a Lua script.

```bash
rbxmk run script.lua
```

If `-` is given, then the script will be read from standard input instead.

```bash
echo 'print("hello world!")' | rbxmk run -
```

The remaining arguments are Lua values to be passed to the file. Numbers, bools,
and nil are parsed into their respective types in Lua, and any other value is
interpreted as a string.

```bash
rbxmk run script.lua true 3.14159 hello!
```

Within the script, these arguments can be received from the `...` operator.

```lua
local arg1, arg2, arg3 = ...
```

The `--desc-*` flags set the `rbxmk.globalDesc` field. They also set the `Enum`
global variable to the enums generated from the descriptor.
