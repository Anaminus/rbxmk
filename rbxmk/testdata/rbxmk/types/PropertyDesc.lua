local desc = rbxmk.newDesc("PropertyDesc")

-- Metamethod tests
T.Pass(typeof(desc) == "PropertyDesc"               , "type of descriptor is PropertyDesc")
T.Pass(type(getmetatable(desc)) == "string"         , "metatable of descriptor is locked")
T.Pass(not string.match(tostring(desc), "^userdata"), "descriptor converts to a string")
T.Pass(desc == desc                                 , "descriptor is equal to itself")
T.Pass(desc ~= rbxmk.newDesc("PropertyDesc")        , "descriptor is not equal to another descriptor of the same type")

-- Member tests

-- ValueType
T.Pass(function() return desc.ValueType end                                                               , "can get ValueType field")
T.Pass(function() return typeof(desc.ValueType) == "TypeDesc" end                                         , "ValueType field is a TypeDesc")
T.Pass(function() return desc.ValueType == rbxmk.newDesc("TypeDesc") end                                  , "ValueType field initializes to empty TypeDesc")
T.Pass(function() desc.ValueType = rbxmk.newDesc("TypeDesc", "FooCategory", "FooName") end                , "can set ValueType field to TypeDesc")
T.Fail(function() desc.ValueType = 42 end                                                                 , "cannot set ValueType field to non-string")
T.Pass(function() return desc.ValueType.Category == "FooCategory" and desc.ValueType.Name == "FooName" end, "set ValueType field persists")

-- ReadSecurity
T.Pass(function() return desc.ReadSecurity end                  , "can get ReadSecurity field")
T.Pass(function() return type(desc.ReadSecurity) == "string" end, "ReadSecurity field is a string")
T.Pass(function() return desc.ReadSecurity == "None" end        , "ReadSecurity field initializes to 'None'")
T.Pass(function() desc.ReadSecurity = "Foobar" end              , "can set ReadSecurity field to string")
T.Fail(function() desc.ReadSecurity = 42 end                    , "cannot set ReadSecurity field to non-string")
T.Pass(function() return desc.ReadSecurity == "Foobar" end      , "set ReadSecurity field persists")

-- WriteSecurity
T.Pass(function() return desc.WriteSecurity end                  , "can get WriteSecurity field")
T.Pass(function() return type(desc.WriteSecurity) == "string" end, "WriteSecurity field is a string")
T.Pass(function() return desc.WriteSecurity == "None" end        , "WriteSecurity field initializes to 'None'")
T.Pass(function() desc.WriteSecurity = "Foobar" end              , "can set WriteSecurity field to string")
T.Fail(function() desc.WriteSecurity = 42 end                    , "cannot set WriteSecurity field to non-string")
T.Pass(function() return desc.WriteSecurity == "Foobar" end      , "set WriteSecurity field persists")

-- CanLoad
T.Pass(function() local _ = desc.CanLoad end                , "can get CanLoad field")
T.Pass(function() return type(desc.CanLoad) == "boolean" end, "CanLoad field is a boolean")
T.Pass(function() return desc.CanLoad == false end          , "CanLoad field initializes to false")
T.Pass(function() desc.CanLoad = true end                   , "can set CanLoad field to boolean")
T.Fail(function() desc.CanLoad = 42 end                     , "cannot set CanLoad field to non-boolean")
T.Pass(function() return desc.CanLoad == true end           , "set CanLoad field persists")

-- CanSave
T.Pass(function() local _ = desc.CanSave end                , "can get CanSave field")
T.Pass(function() return type(desc.CanSave) == "boolean" end, "CanSave field is a boolean")
T.Pass(function() return desc.CanSave == false end          , "CanSave field initializes to false")
T.Pass(function() desc.CanSave = true end                   , "can set CanSave field to boolean")
T.Fail(function() desc.CanSave = 42 end                     , "cannot set CanSave field to non-boolean")
T.Pass(function() return desc.CanSave == true end           , "set CanSave field persists")
