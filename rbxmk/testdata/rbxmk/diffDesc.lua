local prev = RootDesc.new()
local next = RootDesc.new()
T.Fail(function() rbxmk.diffDesc(42, next) end               , "diffDesc expects a RootDesc or nil for its first argument")
T.Fail(function() rbxmk.diffDesc(prev, 42) end               , "diffDesc expects a RootDesc or nil for its second argument")
T.Pass(function() rbxmk.diffDesc(prev, next) end             , "first argument to diffDesc can be a RootDesc")
T.Pass(function() rbxmk.diffDesc(nil, next) end              , "first argument to diffDesc can be nil")
T.Pass(function() rbxmk.diffDesc(prev, next) end             , "second argument to diffDesc can be a RootDesc")
T.Pass(function() rbxmk.diffDesc(prev, nil) end              , "second argument to diffDesc can be nil")
T.Pass(function() rbxmk.diffDesc(nil, nil) end               , "both arguments to diffDesc can be nil")
T.Pass(function() return #rbxmk.diffDesc(prev, next) == 0 end, "diffDesc with no differences returns an empty table")
T.Pass(function() return #rbxmk.diffDesc(nil, nil) == 0 end  , "diffDesc with both nil returns an empty table")

-- TODO: verify correctness of returned actions.
