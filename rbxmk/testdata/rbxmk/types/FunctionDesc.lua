T.Pass(typeof(FunctionDesc.new()) == "FunctionDesc", "new returns FunctionDesc")

local desc = FunctionDesc.new()

-- Metamethod tests
T.Pass(typeof(desc) == "FunctionDesc"               , "type of descriptor is FunctionDesc")
T.Pass(type(getmetatable(desc)) == "string"         , "metatable of descriptor is locked")
T.Pass(not string.match(tostring(desc), "^userdata"), "descriptor converts to a string")
T.Pass(desc == desc                                 , "descriptor is equal to itself")
T.Pass(desc ~= FunctionDesc.new()                   , "descriptor is not equal to another descriptor of the same type")
