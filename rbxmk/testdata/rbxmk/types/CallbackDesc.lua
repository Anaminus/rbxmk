T.Pass(typeof(CallbackDesc.new()) == "CallbackDesc", "new returns CallbackDesc")

local desc = CallbackDesc.new()

-- Metamethod tests
T.Pass(typeof(desc) == "CallbackDesc"               , "type of descriptor is CallbackDesc")
T.Pass(type(getmetatable(desc)) == "string"         , "metatable of descriptor is locked")
T.Pass(not string.match(tostring(desc), "^userdata"), "descriptor converts to a string")
T.Pass(desc == desc                                 , "descriptor is equal to itself")
T.Pass(desc ~= CallbackDesc.new()                   , "descriptor is not equal to another descriptor of the same type")
