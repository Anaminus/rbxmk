local average = os.join(os.expand("$sd"),"average.lua")
T.Pass("runFile receives arguments and returns results",
	rbxmk.runFile(average,1,2,3,4,5,6) == 3.5)
T.Fail("runFile throws error when script errors",
	function() rbxmk.runFile(average,1,2,3,true,5,6) end)
