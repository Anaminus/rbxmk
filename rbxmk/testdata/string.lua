-- Test string.split.
local function checkSplit(str, sep, t)
	local s = string.split(str, sep)
	if #s ~= #t then
		return false
	end
	for i = 1, #t do
		if t[i] ~= s[i] then
			return false
		end
	end
	return true
end

-- Empty string
T.Pass(checkSplit("",nil, {""}))
T.Pass(checkSplit("","|", {""}))

-- Empty substrings
T.Pass(checkSplit("foobar"  , "|" , {"foobar"}))
T.Pass(checkSplit("foo|bar" , "|" , {"foo","bar"}))
T.Pass(checkSplit("foobar"  , ""  , {"f","o","o","b","a","r"}))

T.Pass(checkSplit("foo||bar" , "|" , {"foo","","bar"}))
T.Pass(checkSplit("|foobar"  , "|" , {"","foobar"}))
T.Pass(checkSplit("foobar|"  , "|" , {"foobar",""}))
T.Pass(checkSplit("|"        , "|" , {"",""}))
T.Pass(checkSplit("||"       , "|" , {"","",""}))

T.Pass(checkSplit("foo||bar"   , "||" , {"foo","bar"}))
T.Pass(checkSplit("foo|||bar"  , "||" , {"foo","|bar"}))
T.Pass(checkSplit("foo||||bar" , "||" , {"foo","","bar"}))

-- Whitespace preserved
T.Pass(checkSplit("  whitespace  " , "|" , {"  whitespace  "}))
T.Pass(checkSplit("foo | bar"      , "|" , {"foo ", " bar"}))

-- Invalid UTF-8
T.Pass(checkSplit("\255"      , "|" , {"\255"}))
T.Pass(checkSplit("\253|\254" , "|" , {"\253","\254"}))

-- Unicode
T.Pass(checkSplit("我很高兴，你呢？" , "，" , {"我很高兴","你呢？"}))
T.Pass(checkSplit("hello•world"   , "•"  , {"hello","world"}))

--TODO: strings.Split splits per UTF-8 code point, while string.split splits per
--byte.
--T.Pass(checkSplit("我很高兴，你呢？" , "" , {"\230","\136","\145","\229","\190","\136","\233","\171","\152","\229","\133","\180","\239","\188","\140","\228","\189","\160","\229","\145","\162","\239","\188","\159"}))

--------------------------------------------------------------------------------
--------------------------------------------------------------------------------

T.Pass(function() return string.byte("\255") == 255 end, "test string.byte")
T.Pass(function() return ("\255"):byte() == 255 end, "test string metatables")
