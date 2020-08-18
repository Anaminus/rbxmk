local desc = rbxmk.newDesc("EnumItemDesc")

-- Metamethod tests
T.Pass(typeof(desc) == "EnumItemDesc"               , "type of descriptor is EnumItemDesc")
T.Pass(type(getmetatable(desc)) == "string"         , "metatable of descriptor is locked")
T.Pass(not string.match(tostring(desc), "^userdata"), "descriptor converts to a string")
T.Pass(desc == desc                                 , "descriptor is equal to itself")
T.Pass(desc ~= rbxmk.newDesc("EnumItemDesc")        , "descriptor is not equal to another descriptor of the same type")

-- Value
T.Pass(function() return desc.Value end                  , "can get Value field")
T.Pass(function() return type(desc.Value) == "number" end, "Value field is an int")
T.Pass(function() return desc.Value == 0 end             , "Value field initializes to 0")
T.Pass(function() desc.Value = 42.5 end                  , "can set Value field to int")
T.Fail(function() desc.Value = "Foobar" end              , "cannot set Value field to non-int")
T.Pass(function() return desc.Value == 42 end            , "set Value field persists")

-- Index
T.Pass(function() return desc.Index end                  , "can get Index field")
T.Pass(function() return type(desc.Index) == "number" end, "Index field is an int")
T.Pass(function() return desc.Index == 0 end             , "Index field initializes to 0")
T.Pass(function() desc.Index = 42.5 end                  , "can set Index field to int")
T.Fail(function() desc.Index = "Foobar" end              , "cannot set Index field to non-int")
T.Pass(function() return desc.Index == 42 end            , "set Index field persists")
