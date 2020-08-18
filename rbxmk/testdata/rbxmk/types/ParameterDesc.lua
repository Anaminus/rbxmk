local typeDesc = rbxmk.newDesc("TypeDesc", "FooCategory", "FooName")
local desc = rbxmk.newDesc("ParameterDesc", typeDesc, "fooName")
local descDefault = rbxmk.newDesc("ParameterDesc", typeDesc, "fooName", "FooDefault")

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
