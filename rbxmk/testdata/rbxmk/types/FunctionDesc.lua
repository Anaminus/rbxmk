local desc = rbxmk.newDesc("FunctionDesc")

-- Metamethod tests
T.Pass(typeof(desc) == "FunctionDesc"               , "type of descriptor is FunctionDesc")
T.Pass(type(getmetatable(desc)) == "string"         , "metatable of descriptor is locked")
T.Pass(not string.match(tostring(desc), "^userdata"), "descriptor converts to a string")
T.Pass(desc == desc                                 , "descriptor is equal to itself")
T.Pass(desc ~= rbxmk.newDesc("FunctionDesc")        , "descriptor is not equal to another descriptor of the same type")
