T.Pass(typeof(TypeDesc.new()) == "TypeDesc"          , "new returns TypeDesc")
T.Pass(TypeDesc.new("FooCategory").Category == "FooCategory", "new has optional second string argument that sets Category field")
T.Fail(function() TypeDesc.new(42) end                      , "new with second non-string argument throws an error")
T.Pass(TypeDesc.new(nil, "FooName").Name == "FooName"       , "new has optional third string argument that sets Name field")
T.Fail(function() TypeDesc.new(nil, 42) end                 , "new with third non-string argument throws an error")
T.Pass(function()
	local t = TypeDesc.new("FooCategory", "FooName")
	return t.Category == "FooCategory" and t.Name == "FooName"
end, "new with each argument sets each component")

local desc = TypeDesc.new("FooCategory", "FooName")

-- Metamethod tests
T.Pass(typeof(desc) == "TypeDesc"                    , "type of descriptor is TypeDesc")
T.Pass(type(getmetatable(desc)) == "string"          , "metatable of descriptor is locked")
T.Pass(not string.match(tostring(desc), "^userdata") , "descriptor converts to a string")
T.Pass(desc == desc                                  , "descriptor can be compared with itself")
T.Pass(desc == TypeDesc.new("FooCategory", "FooName"), "descriptor can be compared with a matching TypeDesc")
T.Pass(desc ~= TypeDesc.new("BarCategory", "BarName"), "descriptor can be compared with a non-matching TypeDesc")

-- Member tests

-- Category
T.Fail(function() desc.Category = "Foobar" end             , "descriptor cannot set Category field")
T.Pass(function() return desc.Category == "FooCategory" end, "descriptor can get Category field")

-- Name
T.Fail(function() desc.Name = "Foobar" end         , "descriptor cannot set Name field")
T.Pass(function() return desc.Name == "FooName" end, "descriptor can get Name field")
