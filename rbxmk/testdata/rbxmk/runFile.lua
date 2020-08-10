T.Fail("runFile expects string for its first argument",
	function() rbxmk.runFile(42) end)
T.Fail("runFile throws an error when file does not exist",
	function() rbxmk.runFile("nonextant.lua") end)
T.Pass("runFile receives arguments and returns results",
	rbxmk.runFile(os.expand("$sd/_average.lua"),1,2,3,4,5,6) == 3.5)
T.Fail("runFile throws an error when script errors",
	function() rbxmk.runFile(os.expand("$sd/_average.lua"),1,2,3,true,5,6) end)
T.Fail("runFile throws an error when script is already running",
	function() rbxmk.runFile(os.expand("$sd/$sn")) end)
T.Fail("runFile throws an error when script could not be loaded",
	function() rbxmk.runFile(os.expand("$sd/_badsyntax.lua")) end)
