local root = RootDesc.new()
root:AddClass(ClassDesc.new())
root:AddEnum(EnumDesc.new())
local diff = rbxmk.diffDesc(nil, root)
local desc = diff[1]
local other = diff[2]

-- Metamethod tests
T.Pass(typeof(desc) == "DescAction"                 , "type of value is DescAction")
T.Pass(type(getmetatable(desc)) == "string"         , "metatable of value is locked")
T.Pass(not string.match(tostring(desc), "^userdata"), "value converts to a string")
T.Pass(desc == desc                                 , "value is equal to itself")
T.Pass(desc ~= other                                , "value is not equal to another value of the same type")

-- Constructor tests
T.Fail(function() return DescAction.new() end , "new with no arguments")
T.Fail(function() return DescAction.new(rbxmk.Enum.DescActionType.Add) end , "new with no element")
T.Pass(function() return DescAction.new(rbxmk.Enum.DescActionType.Add, rbxmk.Enum.DescActionElement.Class) end , "new with both arguments")
T.Pass(function() return DescAction.new(1, "Class") end , "new with numeric type and string element")
T.Pass(function() return DescAction.new("Add", 1) end , "new with string type and numeric element")
T.Fail(function() return DescAction.new(100, "Class") end , "new with invalid numeric type")
T.Fail(function() return DescAction.new("INVALID", "Class") end , "new with invalid string type")
T.Fail(function() return DescAction.new("Add", 100) end , "new with invalid numeric element")
T.Fail(function() return DescAction.new("Add", "INVALID") end , "new with invalid string element")

T.Pass(function() return DescAction.new("Remove", "Class").Type.Name == "Remove" end , "new type sets Type")
T.Pass(function() return DescAction.new("Remove", "Class").Element.Name == "Class" end , "new element sets Element")

-- Type property tests
T.Pass(desc.Type == rbxmk.Enum.DescActionType.Add, "desc Type is Add")
T.Pass(function() desc.Type = rbxmk.Enum.DescActionType.Remove end, "set Type to Remove")
T.Pass(desc.Type == rbxmk.Enum.DescActionType.Remove, "desc Type is Remove")
T.Pass(function() desc.Type = 0 end, "set Type to numeric Change")
T.Pass(desc.Type == rbxmk.Enum.DescActionType.Change, "desc Type is Change")
T.Pass(function() desc.Type = "Add" end, "set Type to string Add")
T.Pass(desc.Type == rbxmk.Enum.DescActionType.Add, "desc Type is Add")

-- Element property tests
T.Pass(desc.Element == rbxmk.Enum.DescActionElement.Class, "desc Element is Class")
T.Pass(function() desc.Element = rbxmk.Enum.DescActionElement.Enum end, "set Element to Enum")
T.Pass(desc.Element == rbxmk.Enum.DescActionElement.Enum, "desc Element is Enum")
T.Pass(function() desc.Element = 3 end, "set Element to numeric Function")
T.Pass(desc.Element == rbxmk.Enum.DescActionElement.Function, "desc Element is Function")
T.Pass(function() desc.Element = "EnumItem" end, "set Element to string EnumItem")
T.Pass(desc.Element == rbxmk.Enum.DescActionElement.EnumItem, "desc Element is EnumItem")

-- Primary property tests
T.Pass(desc.Primary == "", "desc Primary is empty")
T.Pass(function() desc.Primary = "Foo" end, "set Primary to Foo")
T.Pass(desc.Primary == "Foo", "desc Primary is Foo")

-- Secondary property tests
T.Pass(desc.Secondary == "", "desc Secondary is empty")
T.Pass(function() desc.Secondary = "Bar" end, "set Secondary to Bar")
T.Pass(desc.Secondary == "Bar", "desc Secondary is Bar")

-- Field tests
T.Pass(desc:Field("Foo") == nil, "Foo field is nil")
T.Pass(desc:Fields().Foo == nil, "desc fields does not contain Foo")

T.Pass(function() desc:SetField("Foo", "Bar") end, "set field Foo to Bar")
T.Pass(desc:Field("Foo") == "Bar", "Foo field is Bar")
T.Pass(desc:Fields().Foo == "Bar", "desc fields contains Foo")

T.Pass(function() desc:SetField("Foo", nil) end, "set field Foo to nil")
T.Pass(desc:Field("Foo") == nil, "Foo field is back to nil")
T.Pass(desc:Fields().Foo == nil, "desc fields no longer contains Foo")

T.Pass(function() desc:SetField("Foo", true) end, "set field Foo to bool")
T.Pass(desc:Field("Foo") == true, "Foo field is bool")
T.Pass(function() desc:SetField("Foo", 42) end, "set field Foo to int")
T.Pass(desc:Field("Foo") == 42, "Foo field is int")
T.Pass(function() desc:SetField("Foo", 42.2) end, "set field Foo to number")
T.Pass(desc:Field("Foo") == 42, "Foo field is number")
T.Pass(function() desc:SetField("Foo", "Bar") end, "set field Foo to string")
T.Pass(desc:Field("Foo") == "Bar", "Foo field is string")
T.Pass(function() desc:SetField("Foo", TypeDesc.new("Foo", "Bar")) end, "set field Foo to TypeDesc")
T.Pass(desc:Field("Foo") == TypeDesc.new("Foo", "Bar"), "Foo field is TypeDesc")
local params = {
	ParameterDesc.new(TypeDesc.new("AA", "BB"), "CC"),
	ParameterDesc.new(TypeDesc.new("DD", "EE"), "FF", "GG"),
	ParameterDesc.new(TypeDesc.new("HH", "II"), "JJ"),
}
T.Fail(function() desc:SetField("Foo", params) end, "cannot set field Foo to parameters")
T.Pass(function() desc:SetField("Parameters", params) end, "set field Parameters to parameters")
T.Pass(#desc:Field("Parameters") == 3, "Parameters field has 3 parameters")
T.Pass(desc:Field("Parameters")[1] == params[1], "Parameters field[1] is param1")
T.Pass(desc:Field("Parameters")[2] == params[2], "Parameters field[2] is param2")
T.Pass(desc:Field("Parameters")[3] == params[3], "Parameters field[3] is param3")
local tags = {"AAAA","BBBB","CCCC"}
T.Fail(function() desc:SetField("Foo", tags) end, "cannot set field Foo to tags")
T.Pass(function() desc:SetField("Tags", tags) end, "set field Tags to tags")
T.Pass(#desc:Field("Tags") == 3, "Tags field has 3 tags")
T.Pass(desc:Field("Tags")[1] == tags[1], "Tags field[1] is tag1")
T.Pass(desc:Field("Tags")[2] == tags[2], "Tags field[2] is tag2")
T.Pass(desc:Field("Tags")[3] == tags[3], "Tags field[3] is tag3")
params[4] = 42
T.Fail(function() desc:SetField("Parameters", params) end, "cannot set field Parameters with non-parameter")
tags[4] = 42
T.Fail(function() desc:SetField("Tags", tags) end, "cannot set field Tags with non-tags")
