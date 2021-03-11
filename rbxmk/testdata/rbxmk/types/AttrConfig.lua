local cfg = rbxmk.newAttrConfig("Foobar")

-- Metamethod tests
T.Pass(typeof(cfg) == "AttrConfig"                  , "type of value")
T.Pass(type(getmetatable(cfg)) == "string"          , "metatable of value is locked")
T.Pass(not string.match(tostring(cfg), "^userdata") , "value converts to a string")
T.Pass(cfg == cfg                                   , "value is equal to itself")
T.Pass(cfg ~= rbxmk.newAttrConfig("Foobar")         , "value is not equal to another value of the same type")
T.Pass(cfg ~= rbxmk.newAttrConfig("Fizzbuzz")       , "value is not equal to another value of the same type, different property")
