T.Pass(typeof(ParameterDesc.new()) == "ParameterDesc", "new returns ParameterDesc")
T.Pass(function()
	local t = TypeDesc.new("FooCategory", "FooName")
	return ParameterDesc.new(t).Type == t
end, "new has optional second TypeDesc argument that sets Type field")
T.Fail(function() ParameterDesc.new(42) end                             , "new with second non-TypeDesc argument throws an error")
T.Pass(ParameterDesc.new(nil, "fooName").Name == "fooName"              , "new has optional third string argument that sets Name field")
T.Fail(function() ParameterDesc.new(nil, 42) end                        , "new with third non-string argument throws an error")
T.Pass(ParameterDesc.new().Default == nil                               , "new has optional fourth argument that, when nil, sets Default to nil")
T.Pass(ParameterDesc.new(nil, nil, "FooDefault").Default == "FooDefault", "new has optional fourth argument that, when a string, sets Default")
T.Fail(function() ParameterDesc.new(nil, nil, 42) end                   , "new with fourth non-string argument throws an error")
T.Pass(function()
	local t = TypeDesc.new("FooCategory", "FooName")
	local p = ParameterDesc.new(t, "fooName", "FooDefault")
	return p.Type == t and p.Name == "fooName" and p.Default == "FooDefault"
end, "new with each argument sets each component")

local typeDesc = TypeDesc.new("FooCategory", "FooName")
local desc = ParameterDesc.new(typeDesc, "fooName")
local descDefault = ParameterDesc.new(typeDesc, "fooName", "FooDefault")

-- Metamethod tests
T.Pass(typeof(desc) == "ParameterDesc"              , "type of descriptor is ParameterDesc")
T.Pass(type(getmetatable(desc)) == "string"         , "metatable of descriptor is locked")
T.Pass(not string.match(tostring(desc), "^userdata"), "descriptor converts to a string")
T.Pass(desc == desc                                 , "descriptor can be compared with itself")
T.Pass(desc == desc                                 , "descriptor can be compared with a matching ParameterDesc")
T.Pass(desc ~= descDefault                          , "descriptor can be compared with a non-matching ParameterDesc")
T.Pass(descDefault == descDefault                   , "descriptor with default can be compared with a matching ParameterDesc")
T.Pass(descDefault ~= desc                          , "descriptor with default can be compared with a non-matching ParameterDesc")

-- Member tests

-- Type
T.Fail(function() desc.Type = typeDesc end        , "descriptor cannot set Type field")
T.Pass(function() return desc.Type == typeDesc end, "descriptor can get Type field")

-- Name
T.Fail(function() desc.Name = "Foobar" end         , "descriptor cannot set Name field")
T.Pass(function() return desc.Name == "fooName" end, "descriptor can get Name field")

-- Default
T.Fail(function() desc.Default = "Foobar" end                   , "descriptor cannot set Default field")
T.Pass(function() return desc.Default == nil end                , "descriptor can get Default field")
T.Fail(function() descDefault.Default = nil end                 , "descriptor with default cannot set Default field")
T.Pass(function() return descDefault.Default == "FooDefault" end, "descriptor with default can get Default field")
