local prev = rbxmk.newDesc("Root")
local next = rbxmk.newDesc("Root")
T.Fail("diffDesc expects a RootDesc or nil for its first argument",
	function() rbxmk.diffDesc(42, next) end)
T.Fail("diffDesc expects a RootDesc or nil for its second argument",
	function() rbxmk.diffDesc(prev, 42) end)
T.Pass("first argument to diffDesc can be a RootDesc",
	function() rbxmk.diffDesc(prev, next) end)
T.Pass("first argument to diffDesc can be nil",
	function() rbxmk.diffDesc(nil, next) end)
T.Pass("second argument to diffDesc can be a RootDesc",
	function() rbxmk.diffDesc(prev, next) end)
T.Pass("second argument to diffDesc can be nil",
	function() rbxmk.diffDesc(prev, nil) end)
T.Pass("both arguments to diffDesc can be nil",
	function() rbxmk.diffDesc(nil, nil) end)
T.Pass("diffDesc with no differences returns an empty table",
	function() return #rbxmk.diffDesc(prev, next) == 0 end)
T.Pass("diffDesc with both nil returns an empty table",
	function() return #rbxmk.diffDesc(nil, nil) == 0 end)

-- TODO: verify correctness of returned actions.
