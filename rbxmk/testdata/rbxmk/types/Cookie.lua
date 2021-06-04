-- new tests
T.Fail(function() return Cookie.new() end                            , "expects a value for first argument")
T.Fail(function() return Cookie.new(42) end                          , "expects a string for its first argument")
T.Fail(function() return Cookie.new("Foo", 42) end                   , "expects a string for its second argument")
T.Pass(Cookie.new("Foo", "Bar")                                      , "can pass string arguments")
T.Pass(typeof(Cookie.new("Foo", "Bar")) == "Cookie"                  , "returns Cookie")
T.Pass(Cookie.new("Foo", "Bar").Name == "Foo"                        , "passing string for first argument sets Name to string")
T.Fail(function() return Cookie.new("Foo", "Bar").Value == "Bar" end , "does not have Value property")

-- from tests
T.Fail(function() return Cookie.from() end , "expects a value for first argument")
T.Fail(function() return Cookie.from("INVALID") end , "expects a known string for first argument")
T.Pass(function() return Cookie.from("studio") or true end , "studio argument succeeds")
T.Pass(function() return Cookie.from("Studio") or true end , "argument is case-insensitive")
local cookies = Cookie.from("studio")
T.Pass(cookies == nil or type(cookies) == "table" , "returns table or nil")
if cookies ~= nil then
	T.Pass(#cookies > 0 , "does not return empty table")
	for i, cookie in ipairs(cookies) do
		T.Pass(typeof(cookie) == "Cookie" , "element " .. i .. " is a Cookie")
	end
end
