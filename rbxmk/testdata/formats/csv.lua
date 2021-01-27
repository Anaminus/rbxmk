local function tableEq(want, got)
	if #want ~= #got then
		return false, string.format("want length %d, got %d", #want, #got)
	end
	for i = 1, #want do
		local w = want[i]
		local g = got[i]
		if type(g) ~= "table" then
			return false, string.format("[%d]: table expected, got %s", i, type(g))
		end
		for j = 1, #w do
			local vw = w[j]
			local vg = g[j]
			if type(vg) ~= "string" then
				return false, string.format("[%d][%d]: string expected, got %s", i, j, type(g))
			end
			if vg ~= vw then
				return false, string.format("[%d][%d]: %q expected, got %q", i, j, vw, vg)
			end
		end
	end
end

local ERROR_FUNC = "(error from function call)"
local function decode(input, contents, msg)
	return function()
		local ok, v = pcall(rbxmk.decodeFormat, "csv", input)
		if not ok then
			if contents == ERROR_FUNC then
				return true
			end
			return false, v
		end
		if type(v) ~= "table" then
			return false, "table expected"
		end
		if contents then
			return tableEq(contents, v)
		end
		return true
	end, msg
end

local function encode(input, contents, msg)
	return function()
		local ok, v = pcall(rbxmk.encodeFormat, "csv", input)
		if not ok then
			if contents == ERROR_FUNC then
				return true
			end
			return false, v
		end
		if contents and v ~= contents then
			return false,
				"\nWANT: " .. (string.gsub(contents, "\n", "\\n")) ..
				"\nGOT : " .. (string.gsub(v, "\n", "\\n"))
		end
		return true
	end, msg
end

T.Pass(decode('A,B,C\nD,E,F\n1,2,3\n', {{"A","B","C"},{"D","E","F"},{"1","2","3"}}), "decode basic")
T.Pass(encode({{"A","B","C"},{"D","E","F"},{"1","2","3"}}, 'A,B,C\nD,E,F\n1,2,3\n'), "encode basic")

T.Pass(decode('A,B,C\nD,E\n1\n', ERROR_FUNC), "decode uneven")
T.Pass(encode({{"A","B","C"},{"D","E"},{"1"}}, 'A,B,C\nD,E\n1\n'), "encode uneven")

T.Pass(decode('', {}), "decode empty")
T.Pass(encode({}, ''), "encode empty")

T.Pass(decode('Key,A,"B",C D,E F,"""G"""\nKK,AA,BB,CD,EF,GG',{{"Key","A","B","C D","E F",'"G"'},{"KK","AA","BB","CD","EF","GG"}}, 'decode csv format'))
T.Pass(encode({{"Key","A","B","C D","E F",'"G"'},{"KK","AA","BB","CD","EF","GG"}}, 'Key,A,B,C D,E F,"""G"""\nKK,AA,BB,CD,EF,GG\n', 'encode csv format'))
