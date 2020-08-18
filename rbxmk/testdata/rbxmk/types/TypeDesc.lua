local desc = rbxmk.newDesc("TypeDesc", "FooCategory", "FooName")

-- Metamethod tests
T.Pass(typeof(desc) == "TypeDesc"                                 , "type of descriptor is TypeDesc")
T.Pass(type(getmetatable(desc)) == "string"                       , "metatable of descriptor is locked")
T.Pass(not string.match(tostring(desc), "^userdata")              , "descriptor converts to a string")
T.Pass(desc == desc                                               , "descriptor can be compared with itself")
T.Pass(desc == rbxmk.newDesc("TypeDesc", "FooCategory", "FooName"), "descriptor can be compared with a matching TypeDesc")
T.Pass(desc ~= rbxmk.newDesc("TypeDesc", "BarCategory", "BarName"), "descriptor can be compared with a non-matching TypeDesc")

-- Member tests

-- Category
T.Fail(function() desc.Category = "Foobar" end             , "descriptor cannot set Category field")
T.Pass(function() return desc.Category == "FooCategory" end, "descriptor can get Category field")

-- Name
T.Fail(function() desc.Name = "Foobar" end         , "descriptor cannot set Name field")
T.Pass(function() return desc.Name == "FooName" end, "descriptor can get Name field")
