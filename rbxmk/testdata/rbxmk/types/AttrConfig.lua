-- Constructor tests
T.Fail(function() AttrConfig.new(42) end                 , "expects a string for its first argument")
T.Pass(AttrConfig.new()                                  , "can pass no value")
T.Pass(AttrConfig.new("Foobar")                          , "can pass string")
T.Pass(typeof(AttrConfig.new("Foobar")) == "AttrConfig"  , "returns AttrConfig")
T.Pass(AttrConfig.new().Property == ""                   , "passing no value sets Property to empty string")
T.Pass(AttrConfig.new("Foobar").Property == "Foobar"     , "passing string sets Property to string")

local cfg = AttrConfig.new("Foobar")

-- Metamethod tests
T.Pass(typeof(cfg) == "AttrConfig"                  , "type of value")
T.Pass(type(getmetatable(cfg)) == "string"          , "metatable of value is locked")
T.Pass(not string.match(tostring(cfg), "^userdata") , "value converts to a string")
T.Pass(cfg == cfg                                   , "value is equal to itself")
T.Pass(cfg ~= AttrConfig.new("Foobar")              , "value is not equal to another value of the same type")
T.Pass(cfg ~= AttrConfig.new("Fizzbuzz")            , "value is not equal to another value of the same type, different property")
