-- meta
local v = Instance.new("BoolValue")
T.Pass("meta sets with 3 arguments",
	select("#", rbxmk.meta(v, "Reference", "foobar")) == 0)
T.Pass("meta gets with 2 arguments",
	select("#", rbxmk.meta(v, "Reference")) == 1)

-- Reference
local v = Instance.new("BoolValue")
T.Pass("Reference initializes with empty string",
	rbxmk.meta(v, "Reference") == "")
T.Pass("Reference is set to string",
	function() rbxmk.meta(v, "Reference", "foobar") end)
T.Fail("Reference errors with non-string",
	function() rbxmk.meta(v, "Reference", true) end)
T.Pass("Reference returns set value",
	rbxmk.meta(v, "Reference") == "foobar")

-- IsService
local v = Instance.new("BoolValue")
T.Pass("IsService initializes with false",
	rbxmk.meta(v, "IsService") == false)
T.Pass("IsService is set to boolean",
	function() rbxmk.meta(v, "IsService", true) end)
T.Fail("IsService errors with non-boolean",
	function() rbxmk.meta(v, "IsService", "foobar") end)
T.Pass("IsService returns set value",
	rbxmk.meta(v, "IsService") == true)

-- Desc
local desc = file.read(os.expand("$sd/../dump.desc.json"))
local v = Instance.new("BoolValue")
T.Pass("Desc initializes with nil",
	rbxmk.meta(v, "Desc") == nil)
T.Pass("RawDesc initializes with nil",
	rbxmk.meta(v, "RawDesc") == nil)
T.Pass("Desc can be set to RootDesc",
	function() rbxmk.meta(v, "Desc", desc) end)
T.Pass("Desc returns RootDesc after being set to RootDesc",
	rbxmk.meta(v, "Desc") == desc)
T.Pass("RawDesc returns RootDesc after Desc is set to RootDesc",
	rbxmk.meta(v, "RawDesc") == desc)
T.Pass("Desc can be set to nil",
	function() rbxmk.meta(v, "Desc", nil) end)
T.Pass("Desc returns nil after being set to nil",
	rbxmk.meta(v, "Desc") == nil)
T.Pass("RawDesc returns nil after Desc is set to nil",
	rbxmk.meta(v, "RawDesc") == nil)
T.Pass("Desc can be set to false",
	function() rbxmk.meta(v, "Desc", false) end)
T.Fail("Desc cannot be set to true",
	function() rbxmk.meta(v, "Desc", true) end)
T.Pass("Desc returns nil after being set to false",
	rbxmk.meta(v, "Desc") == nil)
T.Pass("RawDesc returns false after Desc is set to false",
	rbxmk.meta(v, "RawDesc") == false)
T.Fail("Desc errors without RootDesc, false, or nil",
	function() rbxmk.meta(v, "Desc", "foobar") end)

-- RawDesc
local desc = file.read(os.expand("$sd/../dump.desc.json"))
local v = Instance.new("BoolValue")
T.Pass("RawDesc initializes with nil",
	rbxmk.meta(v, "RawDesc") == nil)
T.Pass("Desc initializes with nil",
	rbxmk.meta(v, "Desc") == nil)
T.Pass("RawDesc can be set to RootDesc",
	function() rbxmk.meta(v, "RawDesc", desc) end)
T.Pass("RawDesc returns RootDesc after being set to RootDesc",
	rbxmk.meta(v, "RawDesc") == desc)
T.Pass("Desc returns RootDesc after RawDesc is set to RootDesc",
	rbxmk.meta(v, "Desc") == desc)
T.Pass("RawDesc can be set to nil",
	function() rbxmk.meta(v, "RawDesc", nil) end)
T.Pass("RawDesc returns nil after being set to nil",
	rbxmk.meta(v, "RawDesc") == nil)
T.Pass("Desc returns nil after RawDesc is set to nil",
	rbxmk.meta(v, "Desc") == nil)
T.Pass("RawDesc can be set to false",
	function() rbxmk.meta(v, "RawDesc", false) end)
T.Fail("RawDesc cannot be set to true",
	function() rbxmk.meta(v, "RawDesc", true) end)
T.Pass("RawDesc returns false after being set to false",
	rbxmk.meta(v, "RawDesc") == false)
T.Pass("Desc returns nil after RawDesc is set to false",
	rbxmk.meta(v, "Desc") == nil)
T.Fail("RawDesc errors without RootDesc, false, or nil",
	function() rbxmk.meta(v, "RawDesc", "foobar") end)

-- Desc inheritance
local a = file.read(os.expand("$sd/../dump.desc.json"))
local b = file.read(os.expand("$sd/../dump.desc.json"))
local n = nil
local f = false

-- Exhaustively check every combination.
local inheritTests = {
	-- [1]: set to Desc of parent
	-- [2]: set to Desc of child
	-- [3]: set to Desc of descendant
	-- [4]: get from Desc of parent
	-- [5]: get from Desc of child
	-- [6]: get from Desc of descendant
	[ 1] = {n,n,n, n,n,n},
	[ 2] = {f,n,n, n,n,n},
	[ 3] = {a,n,n, a,a,a},
	[ 4] = {b,n,n, b,b,b},
	[ 5] = {n,f,n, n,n,n},
	[ 6] = {f,f,n, n,n,n},
	[ 7] = {a,f,n, a,n,n},
	[ 8] = {b,f,n, b,n,n},
	[ 9] = {n,a,n, n,a,a},
	[10] = {f,a,n, n,a,a},
	[11] = {a,a,n, a,a,a},
	[12] = {b,a,n, b,a,a},
	[13] = {n,b,n, n,b,b},
	[14] = {f,b,n, n,b,b},
	[15] = {a,b,n, a,b,b},
	[16] = {b,b,n, b,b,b},
	[17] = {n,n,f, n,n,n},
	[18] = {f,n,f, n,n,n},
	[19] = {a,n,f, a,a,n},
	[20] = {b,n,f, b,b,n},
	[21] = {n,f,f, n,n,n},
	[22] = {f,f,f, n,n,n},
	[23] = {a,f,f, a,n,n},
	[24] = {b,f,f, b,n,n},
	[25] = {n,a,f, n,a,n},
	[26] = {f,a,f, n,a,n},
	[27] = {a,a,f, a,a,n},
	[28] = {b,a,f, b,a,n},
	[29] = {n,b,f, n,b,n},
	[30] = {f,b,f, n,b,n},
	[31] = {a,b,f, a,b,n},
	[32] = {b,b,f, b,b,n},
	[33] = {n,n,a, n,n,a},
	[34] = {f,n,a, n,n,a},
	[35] = {a,n,a, a,a,a},
	[36] = {b,n,a, b,b,a},
	[37] = {n,f,a, n,n,a},
	[38] = {f,f,a, n,n,a},
	[39] = {a,f,a, a,n,a},
	[40] = {b,f,a, b,n,a},
	[41] = {n,a,a, n,a,a},
	[42] = {f,a,a, n,a,a},
	[43] = {a,a,a, a,a,a},
	[44] = {b,a,a, b,a,a},
	[45] = {n,b,a, n,b,a},
	[46] = {f,b,a, n,b,a},
	[47] = {a,b,a, a,b,a},
	[48] = {b,b,a, b,b,a},
	[49] = {n,n,b, n,n,b},
	[50] = {f,n,b, n,n,b},
	[51] = {a,n,b, a,a,b},
	[52] = {b,n,b, b,b,b},
	[53] = {n,f,b, n,n,b},
	[54] = {f,f,b, n,n,b},
	[55] = {a,f,b, a,n,b},
	[56] = {b,f,b, b,n,b},
	[57] = {n,a,b, n,a,b},
	[58] = {f,a,b, n,a,b},
	[59] = {a,a,b, a,a,b},
	[60] = {b,a,b, b,a,b},
	[61] = {n,b,b, n,b,b},
	[62] = {f,b,b, n,b,b},
	[63] = {a,b,b, a,b,b},
	[64] = {b,b,b, b,b,b},
}

local function fmtInheritMessage(i, test, p, c, d)
	local function fmt(v)
		if v == nil then
			return "n"
		elseif v == false then
			return "f"
		elseif v == a then
			return "a"
		elseif v == b then
			return "b"
		else
			return ""
		end
	end
	return string.format("[%2d]: ", i) ..
		fmt(test[1]) .. fmt(test[2]) .. fmt(test[3]) ..
		": want " .. fmt(test[4]) .. fmt(test[5]) .. fmt(test[6]) ..
		", got " .. fmt(p) .. fmt(c) .. fmt(d)
end

local P = Instance.new("BoolValue")
local C = Instance.new("BoolValue", P)
local D = Instance.new("BoolValue", C)
for i, test in ipairs(inheritTests) do
	rbxmk.meta(P, "Desc", test[1])
	rbxmk.meta(C, "Desc", test[2])
	rbxmk.meta(D, "Desc", test[3])
	local p = rbxmk.meta(P, "Desc")
	local c = rbxmk.meta(C, "Desc")
	local d = rbxmk.meta(D, "Desc")
	T.Pass(fmtInheritMessage(i, test, p, c, d), p == test[4] and c == test[5] and d == test[6])
end
