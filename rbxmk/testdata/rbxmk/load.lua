local average = os.join(os.expand("$sd"),"average.lua")
T.Pass("load receives arguments and return results",
	rbxmk.load(average,1,2,3,4,5,6) == 3.5)
T.Fail("load throws error when script errors",
	function() rbxmk.load(average,1,2,3,true,5,6) end)
