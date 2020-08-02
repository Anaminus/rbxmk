local average = os.join(os.expand("$sd"),"average.lua")
PASS(rbxmk.load(average,1,2,3,4,5,6) == 3.5, "receive arguments and return results")
FAIL(function() rbxmk.load(average,1,2,3,true,5,6) end, "throw error when script errors")
