T.Pass(typeof(EventDesc.new()) == "EventDesc", "new returns EventDesc")

local desc = EventDesc.new()

-- Metamethod tests
T.Pass(typeof(desc) == "EventDesc"                  , "type of descriptor is EventDesc")
T.Pass(type(getmetatable(desc)) == "string"         , "metatable of descriptor is locked")
T.Pass(not string.match(tostring(desc), "^userdata"), "descriptor converts to a string")
T.Pass(desc == desc                                 , "descriptor is equal to itself")
T.Pass(desc ~= EventDesc.new()                      , "descriptor is not equal to another descriptor of the same type")
