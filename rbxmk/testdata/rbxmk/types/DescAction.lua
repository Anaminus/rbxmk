local root = rbxmk.newDesc("RootDesc")
root:AddClass(rbxmk.newDesc("ClassDesc"))
root:AddEnum(rbxmk.newDesc("EnumDesc"))
local diff = rbxmk.diffDesc(nil, root)
local desc = diff[1]
local other = diff[2]

-- Metamethod tests
T.Pass("type of value is DescAction",
	typeof(desc) == "DescAction")
T.Pass("metatable of value is locked",
	type(getmetatable(desc)) == "string")
T.Pass("value converts to a string",
	not string.match(tostring(desc), "^userdata"))
T.Pass("value is equal to itself",
	desc == desc)
T.Pass("value is not equal to another value of the same type",
	desc ~= other)
