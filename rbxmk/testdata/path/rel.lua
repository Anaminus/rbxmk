-- https://github.com/golang/go/blob/9cd1818a7d019c02fa4898b3e45a323e35033290/src/path/filepath/path_test.go#L1188-L1240
local tests = {
	-- base        , target           , result
	{ "a/b"        , "a/b"            , "."         },
	{ "a/b/."      , "a/b"            , "."         },
	{ "a/b"        , "a/b/."          , "."         },
	{ "./a/b"      , "a/b"            , "."         },
	{ "a/b"        , "./a/b"          , "."         },
	{ "ab/cd"      , "ab/cde"         , "../cde"    },
	{ "ab/cd"      , "ab/c"           , "../c"      },
	{ "a/b"        , "a/b/c/d"        , "c/d"       },
	{ "a/b"        , "a/b/../c"       , "../c"      },
	{ "a/b/../c"   , "a/b"            , "../b"      },
	{ "a/b/c"      , "a/c/d"          , "../../c/d" },
	{ "a/b"        , "c/d"            , "../../c/d" },
	{ "a/b/c/d"    , "a/b"            , "../.."     },
	{ "a/b/c/d"    , "a/b/"           , "../.."     },
	{ "a/b/c/d/"   , "a/b"            , "../.."     },
	{ "a/b/c/d/"   , "a/b/"           , "../.."     },
	{ "../../a/b"  , "../../a/b/c/d"  , "c/d"       },
	{ "/a/b"       , "/a/b"           , "."         },
	{ "/a/b/."     , "/a/b"           , "."         },
	{ "/a/b"       , "/a/b/."         , "."         },
	{ "/ab/cd"     , "/ab/cde"        , "../cde"    },
	{ "/ab/cd"     , "/ab/c"          , "../c"      },
	{ "/a/b"       , "/a/b/c/d"       , "c/d"       },
	{ "/a/b"       , "/a/b/../c"      , "../c"      },
	{ "/a/b/../c"  , "/a/b"           , "../b"      },
	{ "/a/b/c"     , "/a/c/d"         , "../../c/d" },
	{ "/a/b"       , "/c/d"           , "../../c/d" },
	{ "/a/b/c/d"   , "/a/b"           , "../.."     },
	{ "/a/b/c/d"   , "/a/b/"          , "../.."     },
	{ "/a/b/c/d/"  , "/a/b"           , "../.."     },
	{ "/a/b/c/d/"  , "/a/b/"          , "../.."     },
	{ "/../../a/b" , "/../../a/b/c/d" , "c/d"       },
	{ "."          , "a/b"            , "a/b"       },
	{ "."          , ".."             , ".."        },
	{".."          , "."              , nil         },
	{".."          , "a"              , nil         },
	{"../.."       , ".."             , nil         },
	{"a"           , "/a"             , nil         },
	{"/a"          , "a"              , nil         },

	{win=true, [[C:a\b\c]]          , [[C:a/b/d]]               , [[..\d]]     },
	{win=true, [[C:\]]              , [[D:\]]                   , nil          },
	{win=true, [[C:]]               , [[D:]]                    , nil          },
	{win=true, [[C:\Projects]]      , [[c:\projects\src]]       , [[src]]      },
	{win=true, [[C:\Projects]]      , [[c:\projects]]           , [[.]]        },
	{win=true, [[C:\Projects\a\..]] , [[c:\projects]]           , [[.]]        },
	{win=true, [[\\host\share]]     , [[\\host\share\file.txt]] , [[file.txt]] },
}

local windows = path.clean("/") == "\\"

for i, test in ipairs(tests) do
	if not not test.win == windows then
		local result = path.rel(test[1], test[2])
		if result then
			local expected = path.clean(test[3])
			T.Pass(result == expected, string.format("test %d: expected %q, got %q", i, expected, result))
		else
			T.Pass(result == test[3], "test " .. i .. ": nil result does not match")
		end
	end
end
