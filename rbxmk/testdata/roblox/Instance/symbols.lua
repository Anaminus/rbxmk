local v = Instance.new("BoolValue")
T.Pass(function() v[sym.Reference] = "foobar" end, "can be newindexed with symbols")
T.Pass(function() return v[sym.Reference] end    , "can be indexed with symbols")
T.Fail(function() return v[T.DummySymbol] end    , "can be indexed only with certain symbols")
