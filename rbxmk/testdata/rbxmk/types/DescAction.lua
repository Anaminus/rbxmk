local root = rbxmk.newDesc("RootDesc")
root:AddClass(rbxmk.newDesc("ClassDesc"))
root:AddEnum(rbxmk.newDesc("EnumDesc"))
local diff = rbxmk.diffDesc(nil, root)
local desc = diff[1]
local other = diff[2]

-- Metamethod tests
T.Pass(typeof(desc) == "DescAction"                 , "type of value is DescAction")
T.Pass(type(getmetatable(desc)) == "string"         , "metatable of value is locked")
T.Pass(not string.match(tostring(desc), "^userdata"), "value converts to a string")
T.Pass(desc == desc                                 , "value is equal to itself")
T.Pass(desc ~= other                                , "value is not equal to another value of the same type")
