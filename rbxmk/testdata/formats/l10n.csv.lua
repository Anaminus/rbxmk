local ERROR_FUNC = "(error from function call)"
local function decode(input, contents, msg)
	return function()
		local ok, v = pcall(rbxmk.decodeFormat, "l10n.csv", input)
		if not ok then
			if contents == ERROR_FUNC then
				return true
			end
			return false, v
		end
		if typeof(v) ~= "Instance" then
			return false, "Instance expected"
		end
		if v.ClassName ~= "LocalizationTable" then
			return false, "LocalizationTable expected"
		end
		if contents and v.Contents ~= contents then
			return false,
				"\nWANT: " .. (string.gsub(contents, "\n", "\\n")) ..
				"\nGOT : " .. (string.gsub(v.Contents, "\n", "\\n"))
		end
		return true
	end, msg
end

local function encode(input, contents, msg)
	return function()
		local ok, v = pcall(rbxmk.encodeFormat, "l10n.csv", input)
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

---- Decoding

-- Test empty.
T.Pass(decode(''       , '[]' , "empty csv is valid"))
T.Fail(decode('Header' , '[]' , "empty csv with header expects Key or Source header"))
T.Pass(decode('Key'    , '[]' , "empty csv with Key is valid"))
T.Pass(decode('Source' , '[]' , "empty csv with Source is valid"))

-- Test index headers.
T.Pass(decode('Key\nAA'                                 , '[{"key":"AA","values":{}}]'                                              , 'decode index headers k'))
T.Pass(decode('Source\nAA'                              , '[{"source":"AA","values":{}}]'                                           , 'decode index headers s'))
T.Fail(decode('Context\nAA'                             , '[{"context":"AA","values":{}}]'                                          , 'decode index headers c'))
T.Fail(decode('Example\nAA'                             , '[{"examples":"AA","values":{}}]'                                         , 'decode index headers e'))
T.Pass(decode('Key,Context\nAA,BB'                      , '[{"key":"AA","context":"BB","values":{}}]'                               , 'decode index headers kc'))
T.Pass(decode('Key,Example\nAA,BB'                      , '[{"key":"AA","examples":"BB","values":{}}]'                              , 'decode index headers ke'))
T.Pass(decode('Key,Context,Example,Source\nAA,BB,CC,DD' , '[{"key":"AA","context":"BB","examples":"CC","source":"DD","values":{}}]' , 'decode index headers kces'))

-- Test locale headers.
T.Pass(decode('Key,Foo,bAz,BAR\nAA,BB,CC,DD', '[{"key":"AA","values":{"bar":"DD","baz":"CC","foo":"BB"}}]', 'decode locale headers'))

-- Test repeated headers.
T.Pass(decode('Key,Key\nAA,BB'        , '[{"key":"AA","values":{"key":"BB"}}]' , 'decode repeat headers 0'))
T.Pass(decode('Key,Foo,foo\nAA,BB,CC' , ERROR_FUNC                             , 'decode repeat headers 1'))

-- Test header order.
T.Pass(decode('Foo,Example,Bar,Context,Source,Key\nAA,BB,CC,DD,EE,FF', '[{"key":"FF","context":"DD","examples":"BB","source":"EE","values":{"bar":"CC","foo":"AA"}}]', 'decode header order'))

-- Test multiple records.
T.Pass(decode('Key,Foo\nAA,BB\nCC,DD\nEE,FF', '[{"key":"AA","values":{"foo":"BB"}},{"key":"CC","values":{"foo":"DD"}},{"key":"EE","values":{"foo":"FF"}}]', 'decode multiple records'))

-- Test locale header conflicts.
--
-- Roblox will discard the column if the lowercase header matches the lowercase
-- header of a previous column. If the header is already lowercase, then an
-- error is thrown instead. rbxmk deviates from the Roblox implementation by
-- throwing error for *any* such conflicts.
T.Pass(decode('Key,FOO\nKK,AA\n'               , '[{"key":"KK","values":{"foo":"AA"}}]' , 'decode locale conflict 00'))
T.Pass(decode('Key,Foo\nKK,AA\n'               , '[{"key":"KK","values":{"foo":"AA"}}]' , 'decode locale conflict 01'))
T.Pass(decode('Key,foo\nKK,AA\n'               , '[{"key":"KK","values":{"foo":"AA"}}]' , 'decode locale conflict 02'))
T.Pass(decode('Key,FOO,FOO\nKK,AA,BB\n'        , ERROR_FUNC                             , 'decode locale conflict 03')) -- BB (roblox)
T.Pass(decode('Key,FOO,Foo\nKK,AA,BB\n'        , ERROR_FUNC                             , 'decode locale conflict 04')) -- BB
T.Pass(decode('Key,FOO,foo\nKK,AA,BB\n'        , ERROR_FUNC                             , 'decode locale conflict 05')) -- error
T.Pass(decode('Key,Foo,FOO\nKK,AA,BB\n'        , ERROR_FUNC                             , 'decode locale conflict 06')) -- BB
T.Pass(decode('Key,Foo,Foo\nKK,AA,BB\n'        , ERROR_FUNC                             , 'decode locale conflict 07')) -- BB
T.Pass(decode('Key,Foo,foo\nKK,AA,BB\n'        , ERROR_FUNC                             , 'decode locale conflict 08')) -- error
T.Pass(decode('Key,foo,FOO\nKK,AA,BB\n'        , ERROR_FUNC                             , 'decode locale conflict 09')) -- BB
T.Pass(decode('Key,foo,Foo\nKK,AA,BB\n'        , ERROR_FUNC                             , 'decode locale conflict 10')) -- BB
T.Pass(decode('Key,foo,foo\nKK,AA,BB\n'        , ERROR_FUNC                             , 'decode locale conflict 11')) -- error
T.Pass(decode('Key,FOO,FOO,FOO\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 12')) -- CC
T.Pass(decode('Key,Foo,FOO,FOO\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 13')) -- CC
T.Pass(decode('Key,foo,FOO,FOO\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 14')) -- CC
T.Pass(decode('Key,FOO,Foo,FOO\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 15')) -- CC
T.Pass(decode('Key,Foo,Foo,FOO\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 16')) -- CC
T.Pass(decode('Key,foo,Foo,FOO\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 17')) -- CC
T.Pass(decode('Key,FOO,foo,FOO\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 18')) -- error
T.Pass(decode('Key,Foo,foo,FOO\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 19')) -- error
T.Pass(decode('Key,foo,foo,FOO\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 20')) -- error
T.Pass(decode('Key,FOO,FOO,Foo\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 21')) -- CC
T.Pass(decode('Key,Foo,FOO,Foo\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 22')) -- CC
T.Pass(decode('Key,foo,FOO,Foo\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 23')) -- CC
T.Pass(decode('Key,FOO,Foo,Foo\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 24')) -- CC
T.Pass(decode('Key,Foo,Foo,Foo\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 25')) -- CC
T.Pass(decode('Key,foo,Foo,Foo\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 26')) -- CC
T.Pass(decode('Key,FOO,foo,Foo\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 27')) -- error
T.Pass(decode('Key,Foo,foo,Foo\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 28')) -- error
T.Pass(decode('Key,foo,foo,Foo\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 29')) -- error
T.Pass(decode('Key,FOO,FOO,foo\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 30')) -- error
T.Pass(decode('Key,Foo,FOO,foo\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 31')) -- error
T.Pass(decode('Key,foo,FOO,foo\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 32')) -- error
T.Pass(decode('Key,FOO,Foo,foo\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 33')) -- error
T.Pass(decode('Key,Foo,Foo,foo\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 34')) -- error
T.Pass(decode('Key,foo,Foo,foo\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 35')) -- error
T.Pass(decode('Key,FOO,foo,foo\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 36')) -- error
T.Pass(decode('Key,Foo,foo,foo\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 37')) -- error
T.Pass(decode('Key,foo,foo,foo\nKK,AA,BB,CC\n' , ERROR_FUNC                             , 'decode locale conflict 38')) -- error

-- Test index header conflicts.
T.Pass(decode('Key,Context,Source,foo\n,,,A\n,,,B\n'       , ERROR_FUNC                                                                                                        , 'decode index conflict 00'))
T.Pass(decode('Key,Context,Source,foo\nK,,,A\n,,,B\n'      , ERROR_FUNC                                                                                                        , 'decode index conflict 01'))
T.Pass(decode('Key,Context,Source,foo\n,C,,A\n,,,B\n'      , ERROR_FUNC                                                                                                        , 'decode index conflict 02'))
T.Pass(decode('Key,Context,Source,foo\n,,S,A\n,,,B\n'      , ERROR_FUNC                                                                                                        , 'decode index conflict 03'))
T.Pass(decode('Key,Context,Source,foo\nK,C,,A\n,,,B\n'     , ERROR_FUNC                                                                                                        , 'decode index conflict 04'))
T.Pass(decode('Key,Context,Source,foo\nK,,S,A\n,,,B\n'     , ERROR_FUNC                                                                                                        , 'decode index conflict 05'))
T.Pass(decode('Key,Context,Source,foo\n,C,S,A\n,,,B\n'     , ERROR_FUNC                                                                                                        , 'decode index conflict 06'))
T.Pass(decode('Key,Context,Source,foo\nK,C,S,A\n,,,B\n'    , ERROR_FUNC                                                                                                        , 'decode index conflict 07'))
T.Pass(decode('Key,Context,Source,foo\n,,,A\nK,,,B\n'      , ERROR_FUNC                                                                                                        , 'decode index conflict 08'))
T.Pass(decode('Key,Context,Source,foo\nK,,,A\nK,,,B\n'     , ERROR_FUNC                                                                                                        , 'decode index conflict 09'))
T.Pass(decode('Key,Context,Source,foo\n,C,,A\nK,,,B\n'     , ERROR_FUNC                                                                                                        , 'decode index conflict 10'))
T.Pass(decode('Key,Context,Source,foo\n,,S,A\nK,,,B\n'     , '[{"source":"S","values":{"foo":"A"}},{"key":"K","values":{"foo":"B"}}]'                                          , 'decode index conflict 11'))
T.Pass(decode('Key,Context,Source,foo\nK,C,,A\nK,,,B\n'    , ERROR_FUNC                                                                                                        , 'decode index conflict 12'))
T.Pass(decode('Key,Context,Source,foo\nK,,S,A\nK,,,B\n'    , ERROR_FUNC                                                                                                        , 'decode index conflict 13'))
T.Pass(decode('Key,Context,Source,foo\n,C,S,A\nK,,,B\n'    , '[{"context":"C","source":"S","values":{"foo":"A"}},{"key":"K","values":{"foo":"B"}}]'                            , 'decode index conflict 14'))
T.Pass(decode('Key,Context,Source,foo\nK,C,S,A\nK,,,B\n'   , ERROR_FUNC                                                                                                        , 'decode index conflict 15'))
T.Pass(decode('Key,Context,Source,foo\n,,,A\n,C,,B\n'      , ERROR_FUNC                                                                                                        , 'decode index conflict 16'))
T.Pass(decode('Key,Context,Source,foo\nK,,,A\n,C,,B\n'     , ERROR_FUNC                                                                                                        , 'decode index conflict 17'))
T.Pass(decode('Key,Context,Source,foo\n,C,,A\n,C,,B\n'     , ERROR_FUNC                                                                                                        , 'decode index conflict 18'))
T.Pass(decode('Key,Context,Source,foo\n,,S,A\n,C,,B\n'     , ERROR_FUNC                                                                                                        , 'decode index conflict 19'))
T.Pass(decode('Key,Context,Source,foo\nK,C,,A\n,C,,B\n'    , ERROR_FUNC                                                                                                        , 'decode index conflict 20'))
T.Pass(decode('Key,Context,Source,foo\nK,,S,A\n,C,,B\n'    , ERROR_FUNC                                                                                                        , 'decode index conflict 21'))
T.Pass(decode('Key,Context,Source,foo\n,C,S,A\n,C,,B\n'    , ERROR_FUNC                                                                                                        , 'decode index conflict 22'))
T.Pass(decode('Key,Context,Source,foo\nK,C,S,A\n,C,,B\n'   , ERROR_FUNC                                                                                                        , 'decode index conflict 23'))
T.Pass(decode('Key,Context,Source,foo\n,,,A\n,,S,B\n'      , ERROR_FUNC                                                                                                        , 'decode index conflict 24'))
T.Pass(decode('Key,Context,Source,foo\nK,,,A\n,,S,B\n'     , '[{"key":"K","values":{"foo":"A"}},{"source":"S","values":{"foo":"B"}}]'                                          , 'decode index conflict 25'))
T.Pass(decode('Key,Context,Source,foo\n,C,,A\n,,S,B\n'     , ERROR_FUNC                                                                                                        , 'decode index conflict 26'))
T.Pass(decode('Key,Context,Source,foo\n,,S,A\n,,S,B\n'     , ERROR_FUNC                                                                                                        , 'decode index conflict 27'))
T.Pass(decode('Key,Context,Source,foo\nK,C,,A\n,,S,B\n'    , '[{"key":"K","context":"C","values":{"foo":"A"}},{"source":"S","values":{"foo":"B"}}]'                            , 'decode index conflict 28'))
T.Pass(decode('Key,Context,Source,foo\nK,,S,A\n,,S,B\n'    , '[{"key":"K","source":"S","values":{"foo":"A"}},{"source":"S","values":{"foo":"B"}}]'                             , 'decode index conflict 29'))
T.Pass(decode('Key,Context,Source,foo\n,C,S,A\n,,S,B\n'    , '[{"context":"C","source":"S","values":{"foo":"A"}},{"source":"S","values":{"foo":"B"}}]'                         , 'decode index conflict 30'))
T.Pass(decode('Key,Context,Source,foo\nK,C,S,A\n,,S,B\n'   , '[{"key":"K","context":"C","source":"S","values":{"foo":"A"}},{"source":"S","values":{"foo":"B"}}]'               , 'decode index conflict 31'))
T.Pass(decode('Key,Context,Source,foo\n,,,A\nK,C,,B\n'     , ERROR_FUNC                                                                                                        , 'decode index conflict 32'))
T.Pass(decode('Key,Context,Source,foo\nK,,,A\nK,C,,B\n'    , ERROR_FUNC                                                                                                        , 'decode index conflict 33'))
T.Pass(decode('Key,Context,Source,foo\n,C,,A\nK,C,,B\n'    , ERROR_FUNC                                                                                                        , 'decode index conflict 34'))
T.Pass(decode('Key,Context,Source,foo\n,,S,A\nK,C,,B\n'    , '[{"source":"S","values":{"foo":"A"}},{"key":"K","context":"C","values":{"foo":"B"}}]'                            , 'decode index conflict 35'))
T.Pass(decode('Key,Context,Source,foo\nK,C,,A\nK,C,,B\n'   , ERROR_FUNC                                                                                                        , 'decode index conflict 36'))
T.Pass(decode('Key,Context,Source,foo\nK,,S,A\nK,C,,B\n'   , ERROR_FUNC                                                                                                        , 'decode index conflict 37'))
T.Pass(decode('Key,Context,Source,foo\n,C,S,A\nK,C,,B\n'   , '[{"context":"C","source":"S","values":{"foo":"A"}},{"key":"K","context":"C","values":{"foo":"B"}}]'              , 'decode index conflict 38'))
T.Pass(decode('Key,Context,Source,foo\nK,C,S,A\nK,C,,B\n'  , ERROR_FUNC                                                                                                        , 'decode index conflict 39'))
T.Pass(decode('Key,Context,Source,foo\n,,,A\nK,,S,B\n'     , ERROR_FUNC                                                                                                        , 'decode index conflict 40'))
T.Pass(decode('Key,Context,Source,foo\nK,,,A\nK,,S,B\n'    , ERROR_FUNC                                                                                                        , 'decode index conflict 41'))
T.Pass(decode('Key,Context,Source,foo\n,C,,A\nK,,S,B\n'    , ERROR_FUNC                                                                                                        , 'decode index conflict 42'))
T.Pass(decode('Key,Context,Source,foo\n,,S,A\nK,,S,B\n'    , '[{"source":"S","values":{"foo":"A"}},{"key":"K","source":"S","values":{"foo":"B"}}]'                             , 'decode index conflict 43'))
T.Pass(decode('Key,Context,Source,foo\nK,C,,A\nK,,S,B\n'   , ERROR_FUNC                                                                                                        , 'decode index conflict 44'))
T.Pass(decode('Key,Context,Source,foo\nK,,S,A\nK,,S,B\n'   , ERROR_FUNC                                                                                                        , 'decode index conflict 45'))
T.Pass(decode('Key,Context,Source,foo\n,C,S,A\nK,,S,B\n'   , '[{"context":"C","source":"S","values":{"foo":"A"}},{"key":"K","source":"S","values":{"foo":"B"}}]'               , 'decode index conflict 46'))
T.Pass(decode('Key,Context,Source,foo\nK,C,S,A\nK,,S,B\n'  , ERROR_FUNC                                                                                                        , 'decode index conflict 47'))
T.Pass(decode('Key,Context,Source,foo\n,,,A\n,C,S,B\n'     , ERROR_FUNC                                                                                                        , 'decode index conflict 48'))
T.Pass(decode('Key,Context,Source,foo\nK,,,A\n,C,S,B\n'    , '[{"key":"K","values":{"foo":"A"}},{"context":"C","source":"S","values":{"foo":"B"}}]'                            , 'decode index conflict 49'))
T.Pass(decode('Key,Context,Source,foo\n,C,,A\n,C,S,B\n'    , ERROR_FUNC                                                                                                        , 'decode index conflict 50'))
T.Pass(decode('Key,Context,Source,foo\n,,S,A\n,C,S,B\n'    , '[{"source":"S","values":{"foo":"A"}},{"context":"C","source":"S","values":{"foo":"B"}}]'                         , 'decode index conflict 51'))
T.Pass(decode('Key,Context,Source,foo\nK,C,,A\n,C,S,B\n'   , '[{"key":"K","context":"C","values":{"foo":"A"}},{"context":"C","source":"S","values":{"foo":"B"}}]'              , 'decode index conflict 52'))
T.Pass(decode('Key,Context,Source,foo\nK,,S,A\n,C,S,B\n'   , '[{"key":"K","source":"S","values":{"foo":"A"}},{"context":"C","source":"S","values":{"foo":"B"}}]'               , 'decode index conflict 53'))
T.Pass(decode('Key,Context,Source,foo\n,C,S,A\n,C,S,B\n'   , ERROR_FUNC                                                                                                        , 'decode index conflict 54'))
T.Pass(decode('Key,Context,Source,foo\nK,C,S,A\n,C,S,B\n'  , '[{"key":"K","context":"C","source":"S","values":{"foo":"A"}},{"context":"C","source":"S","values":{"foo":"B"}}]' , 'decode index conflict 55'))
T.Pass(decode('Key,Context,Source,foo\n,,,A\nK,C,S,B\n'    , ERROR_FUNC                                                                                                        , 'decode index conflict 56'))
T.Pass(decode('Key,Context,Source,foo\nK,,,A\nK,C,S,B\n'   , ERROR_FUNC                                                                                                        , 'decode index conflict 57'))
T.Pass(decode('Key,Context,Source,foo\n,C,,A\nK,C,S,B\n'   , ERROR_FUNC                                                                                                        , 'decode index conflict 58'))
T.Pass(decode('Key,Context,Source,foo\n,,S,A\nK,C,S,B\n'   , '[{"source":"S","values":{"foo":"A"}},{"key":"K","context":"C","source":"S","values":{"foo":"B"}}]'               , 'decode index conflict 59'))
T.Pass(decode('Key,Context,Source,foo\nK,C,,A\nK,C,S,B\n'  , ERROR_FUNC                                                                                                        , 'decode index conflict 60'))
T.Pass(decode('Key,Context,Source,foo\nK,,S,A\nK,C,S,B\n'  , ERROR_FUNC                                                                                                        , 'decode index conflict 61'))
T.Pass(decode('Key,Context,Source,foo\n,C,S,A\nK,C,S,B\n'  , '[{"context":"C","source":"S","values":{"foo":"A"}},{"key":"K","context":"C","source":"S","values":{"foo":"B"}}]' , 'decode index conflict 62'))
T.Pass(decode('Key,Context,Source,foo\nK,C,S,A\nK,C,S,B\n' , ERROR_FUNC                                                                                                        , 'decode index conflict 63'))

-- Test unmerged headers.
T.Pass(decode('Key,a,b,c\n11,AA,BB,\n22,,BB,CC\n', '[{"key":"11","values":{"a":"AA","b":"BB","c":""}},{"key":"22","values":{"a":"","b":"BB","c":"CC"}}]', 'decode merge headers'))

-- Test CSV format.
T.Pass(decode('Key,A,"B",C D,E F,"""G"""\nKK,AA,BB,CD,EF,GG', '[{"key":"KK","values":{"\\"g\\"":"GG","a":"AA","b":"BB","c d":"CD","e f":"EF"}}]', 'decode csv format'))

--------------------------------------------------------------------------------
--------------------------------------------------------------------------------
-- Encoding

-- Test empty.
T.Pass(encode('[]', 'Key,Context,Example,Source\n', "empty csv is valid"))

-- Test index headers.
T.Pass(encode('[{"key":"AA","values":{}}]'                                              , 'Key,Context,Example,Source\nAA,,,\n'       , 'encode index headers k'))
T.Pass(encode('[{"source":"AA","values":{}}]'                                           , 'Key,Context,Example,Source\n,,,AA\n'       , 'encode index headers s'))
T.Pass(encode('[{"key":"AA","context":"BB","values":{}}]'                               , 'Key,Context,Example,Source\nAA,BB,,\n'     , 'encode index headers kc'))
T.Pass(encode('[{"key":"AA","examples":"BB","values":{}}]'                              , 'Key,Context,Example,Source\nAA,,BB,\n'     , 'encode index headers ke'))
T.Pass(encode('[{"key":"AA","context":"BB","examples":"CC","source":"DD","values":{}}]' , 'Key,Context,Example,Source\nAA,BB,CC,DD\n' , 'encode index headers kces'))

-- Test locale headers.
T.Pass(encode('[{"key":"AA","values":{"bar":"DD","baz":"CC","foo":"BB"}}]', 'Key,Context,Example,Source,bar,baz,foo\nAA,,,,DD,CC,BB\n', 'encode locale headers'))

-- Test repeated headers.
T.Pass(encode('[{"key":"AA","values":{"key":"BB"}}]', 'Key,Context,Example,Source,key\nAA,,,,BB\n', 'encode repeat headers'))

-- Test header order.
T.Pass(encode('[{"key":"FF","context":"DD","examples":"BB","source":"EE","values":{"bar":"CC","foo":"AA"}}]', 'Key,Context,Example,Source,bar,foo\nFF,DD,BB,EE,CC,AA\n', 'encode header order'))

-- Test multiple records.
T.Pass(encode('[{"key":"AA","values":{"foo":"BB"}},{"key":"CC","values":{"foo":"DD"}},{"key":"EE","values":{"foo":"FF"}}]', 'Key,Context,Example,Source,foo\nAA,,,,BB\nCC,,,,DD\nEE,,,,FF\n', 'encode multiple records'))

-- Test index header conflicts.
T.Pass(encode('[{"key":"KK","values":{"FOO":"AA"}}]'                       , 'Key,Context,Example,Source,foo\nKK,,,,AA\n'               , 'encode header conflict 00')) -- AA (roblox)
T.Pass(encode('[{"key":"KK","values":{"Foo":"AA"}}]'                       , 'Key,Context,Example,Source,foo\nKK,,,,AA\n'               , 'encode header conflict 01')) -- AA
T.Pass(encode('[{"key":"KK","values":{"foo":"AA"}}]'                       , 'Key,Context,Example,Source,foo\nKK,,,,AA\n'               , 'encode header conflict 02')) -- AA
T.Pass(encode('[{"key":"KK","values":{"FOO":"AA","FOO":"BB"}}]'            , 'Key,Context,Example,Source,foo\nKK,,,,BB\n'               , 'encode header conflict 03')) -- BB
T.Pass(encode('[{"key":"KK","values":{"FOO":"AA","Foo":"BB"}}]'            , ERROR_FUNC                                                 , 'encode header conflict 04')) -- BB
T.Pass(encode('[{"key":"KK","values":{"FOO":"AA","foo":"BB"}}]'            , ERROR_FUNC                                                 , 'encode header conflict 05')) -- BB
T.Pass(encode('[{"key":"KK","values":{"Foo":"AA","FOO":"BB"}}]'            , ERROR_FUNC                                                 , 'encode header conflict 06')) -- BB
T.Pass(encode('[{"key":"KK","values":{"Foo":"AA","Foo":"BB"}}]'            , 'Key,Context,Example,Source,foo\nKK,,,,BB\n'               , 'encode header conflict 07')) -- BB
T.Pass(encode('[{"key":"KK","values":{"Foo":"AA","foo":"BB"}}]'            , ERROR_FUNC                                                 , 'encode header conflict 08')) -- BB
T.Pass(encode('[{"key":"KK","values":{"foo":"AA","FOO":"BB"}}]'            , ERROR_FUNC                                                 , 'encode header conflict 09')) -- BB
T.Pass(encode('[{"key":"KK","values":{"foo":"AA","Foo":"BB"}}]'            , ERROR_FUNC                                                 , 'encode header conflict 10')) -- BB
T.Pass(encode('[{"key":"KK","values":{"foo":"AA","foo":"BB"}}]'            , 'Key,Context,Example,Source,foo\nKK,,,,BB\n'               , 'encode header conflict 11')) -- BB
T.Pass(encode('[{"key":"KK","values":{"FOO":"AA","FOO":"BB","FOO":"CC"}}]' , 'Key,Context,Example,Source,foo\nKK,,,,CC\n'               , 'encode header conflict 12')) -- CC
T.Pass(encode('[{"key":"KK","values":{"Foo":"AA","FOO":"BB","FOO":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 13')) -- CC
T.Pass(encode('[{"key":"KK","values":{"foo":"AA","FOO":"BB","FOO":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 14')) -- CC
T.Pass(encode('[{"key":"KK","values":{"FOO":"AA","Foo":"BB","FOO":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 15')) -- CC
T.Pass(encode('[{"key":"KK","values":{"Foo":"AA","Foo":"BB","FOO":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 16')) -- CC
T.Pass(encode('[{"key":"KK","values":{"foo":"AA","Foo":"BB","FOO":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 17')) -- CC
T.Pass(encode('[{"key":"KK","values":{"FOO":"AA","foo":"BB","FOO":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 18')) -- CC
T.Pass(encode('[{"key":"KK","values":{"Foo":"AA","foo":"BB","FOO":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 19')) -- CC
T.Pass(encode('[{"key":"KK","values":{"foo":"AA","foo":"BB","FOO":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 20')) -- CC
T.Pass(encode('[{"key":"KK","values":{"FOO":"AA","FOO":"BB","Foo":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 21')) -- CC
T.Pass(encode('[{"key":"KK","values":{"Foo":"AA","FOO":"BB","Foo":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 22')) -- CC
T.Pass(encode('[{"key":"KK","values":{"foo":"AA","FOO":"BB","Foo":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 23')) -- CC
T.Pass(encode('[{"key":"KK","values":{"FOO":"AA","Foo":"BB","Foo":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 24')) -- CC
T.Pass(encode('[{"key":"KK","values":{"Foo":"AA","Foo":"BB","Foo":"CC"}}]' , 'Key,Context,Example,Source,foo\nKK,,,,CC\n'               , 'encode header conflict 25')) -- CC
T.Pass(encode('[{"key":"KK","values":{"foo":"AA","Foo":"BB","Foo":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 26')) -- CC
T.Pass(encode('[{"key":"KK","values":{"FOO":"AA","foo":"BB","Foo":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 27')) -- CC
T.Pass(encode('[{"key":"KK","values":{"Foo":"AA","foo":"BB","Foo":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 28')) -- CC
T.Pass(encode('[{"key":"KK","values":{"foo":"AA","foo":"BB","Foo":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 29')) -- CC
T.Pass(encode('[{"key":"KK","values":{"FOO":"AA","FOO":"BB","foo":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 30')) -- CC
T.Pass(encode('[{"key":"KK","values":{"Foo":"AA","FOO":"BB","foo":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 31')) -- CC
T.Pass(encode('[{"key":"KK","values":{"foo":"AA","FOO":"BB","foo":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 32')) -- CC
T.Pass(encode('[{"key":"KK","values":{"FOO":"AA","Foo":"BB","foo":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 33')) -- CC
T.Pass(encode('[{"key":"KK","values":{"Foo":"AA","Foo":"BB","foo":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 34')) -- CC
T.Pass(encode('[{"key":"KK","values":{"foo":"AA","Foo":"BB","foo":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 35')) -- CC
T.Pass(encode('[{"key":"KK","values":{"FOO":"AA","foo":"BB","foo":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 36')) -- CC
T.Pass(encode('[{"key":"KK","values":{"Foo":"AA","foo":"BB","foo":"CC"}}]' , ERROR_FUNC                                                 , 'encode header conflict 37')) -- CC
T.Pass(encode('[{"key":"KK","values":{"foo":"AA","foo":"BB","foo":"CC"}}]' , 'Key,Context,Example,Source,foo\nKK,,,,CC\n'               , 'encode header conflict 38')) -- CC

-- Test conflicting index headers.
T.Pass(encode('[{"values":{"foo":"A"}},{"values":{"foo":"B"}}]'                                                                           , ERROR_FUNC                                            , 'encode index conflict 00'))
T.Pass(encode('[{"key":"K","values":{"foo":"A"}},{"values":{"foo":"B"}}]'                                                                 , ERROR_FUNC                                            , 'encode index conflict 01'))
T.Pass(encode('[{"context":"C","values":{"foo":"A"}},{"values":{"foo":"B"}}]'                                                             , ERROR_FUNC                                            , 'encode index conflict 02'))
T.Pass(encode('[{"source":"S","values":{"foo":"A"}},{"values":{"foo":"B"}}]'                                                              , ERROR_FUNC                                            , 'encode index conflict 03'))
T.Pass(encode('[{"key":"K","context":"C","values":{"foo":"A"}},{"values":{"foo":"B"}}]'                                                   , ERROR_FUNC                                            , 'encode index conflict 04'))
T.Pass(encode('[{"key":"K","source":"S","values":{"foo":"A"}},{"values":{"foo":"B"}}]'                                                    , ERROR_FUNC                                            , 'encode index conflict 05'))
T.Pass(encode('[{"context":"C","source":"S","values":{"foo":"A"}},{"values":{"foo":"B"}}]'                                                , ERROR_FUNC                                            , 'encode index conflict 06'))
T.Pass(encode('[{"key":"K","context":"C","source":"S","values":{"foo":"A"}},{"values":{"foo":"B"}}]'                                      , ERROR_FUNC                                            , 'encode index conflict 07'))
T.Pass(encode('[{"values":{"foo":"A"}},{"key":"K","values":{"foo":"B"}}]'                                                                 , ERROR_FUNC                                            , 'encode index conflict 08'))
T.Pass(encode('[{"key":"K","values":{"foo":"A"}},{"key":"K","values":{"foo":"B"}}]'                                                       , ERROR_FUNC                                            , 'encode index conflict 09'))
T.Pass(encode('[{"context":"C","values":{"foo":"A"}},{"key":"K","values":{"foo":"B"}}]'                                                   , ERROR_FUNC                                            , 'encode index conflict 10'))
T.Pass(encode('[{"source":"S","values":{"foo":"A"}},{"key":"K","values":{"foo":"B"}}]'                                                    , 'Key,Context,Example,Source,foo\n,,,S,A\nK,,,,B\n'    , 'encode index conflict 11'))
T.Pass(encode('[{"key":"K","context":"C","values":{"foo":"A"}},{"key":"K","values":{"foo":"B"}}]'                                         , ERROR_FUNC                                            , 'encode index conflict 12'))
T.Pass(encode('[{"key":"K","source":"S","values":{"foo":"A"}},{"key":"K","values":{"foo":"B"}}]'                                          , ERROR_FUNC                                            , 'encode index conflict 13'))
T.Pass(encode('[{"context":"C","source":"S","values":{"foo":"A"}},{"key":"K","values":{"foo":"B"}}]'                                      , 'Key,Context,Example,Source,foo\n,C,,S,A\nK,,,,B\n'   , 'encode index conflict 14'))
T.Pass(encode('[{"key":"K","context":"C","source":"S","values":{"foo":"A"}},{"key":"K","values":{"foo":"B"}}]'                            , ERROR_FUNC                                            , 'encode index conflict 15'))
T.Pass(encode('[{"values":{"foo":"A"}},{"context":"C","values":{"foo":"B"}}]'                                                             , ERROR_FUNC                                            , 'encode index conflict 16'))
T.Pass(encode('[{"key":"K","values":{"foo":"A"}},{"context":"C","values":{"foo":"B"}}]'                                                   , ERROR_FUNC                                            , 'encode index conflict 17'))
T.Pass(encode('[{"context":"C","values":{"foo":"A"}},{"context":"C","values":{"foo":"B"}}]'                                               , ERROR_FUNC                                            , 'encode index conflict 18'))
T.Pass(encode('[{"source":"S","values":{"foo":"A"}},{"context":"C","values":{"foo":"B"}}]'                                                , ERROR_FUNC                                            , 'encode index conflict 19'))
T.Pass(encode('[{"key":"K","context":"C","values":{"foo":"A"}},{"context":"C","values":{"foo":"B"}}]'                                     , ERROR_FUNC                                            , 'encode index conflict 20'))
T.Pass(encode('[{"key":"K","source":"S","values":{"foo":"A"}},{"context":"C","values":{"foo":"B"}}]'                                      , ERROR_FUNC                                            , 'encode index conflict 21'))
T.Pass(encode('[{"context":"C","source":"S","values":{"foo":"A"}},{"context":"C","values":{"foo":"B"}}]'                                  , ERROR_FUNC                                            , 'encode index conflict 22'))
T.Pass(encode('[{"key":"K","context":"C","source":"S","values":{"foo":"A"}},{"context":"C","values":{"foo":"B"}}]'                        , ERROR_FUNC                                            , 'encode index conflict 23'))
T.Pass(encode('[{"values":{"foo":"A"}},{"source":"S","values":{"foo":"B"}}]'                                                              , ERROR_FUNC                                            , 'encode index conflict 24'))
T.Pass(encode('[{"key":"K","values":{"foo":"A"}},{"source":"S","values":{"foo":"B"}}]'                                                    , 'Key,Context,Example,Source,foo\nK,,,,A\n,,,S,B\n'    , 'encode index conflict 25'))
T.Pass(encode('[{"context":"C","values":{"foo":"A"}},{"source":"S","values":{"foo":"B"}}]'                                                , ERROR_FUNC                                            , 'encode index conflict 26'))
T.Pass(encode('[{"source":"S","values":{"foo":"A"}},{"source":"S","values":{"foo":"B"}}]'                                                 , ERROR_FUNC                                            , 'encode index conflict 27'))
T.Pass(encode('[{"key":"K","context":"C","values":{"foo":"A"}},{"source":"S","values":{"foo":"B"}}]'                                      , 'Key,Context,Example,Source,foo\nK,C,,,A\n,,,S,B\n'   , 'encode index conflict 28'))
T.Pass(encode('[{"key":"K","source":"S","values":{"foo":"A"}},{"source":"S","values":{"foo":"B"}}]'                                       , 'Key,Context,Example,Source,foo\nK,,,S,A\n,,,S,B\n'   , 'encode index conflict 29'))
T.Pass(encode('[{"context":"C","source":"S","values":{"foo":"A"}},{"source":"S","values":{"foo":"B"}}]'                                   , 'Key,Context,Example,Source,foo\n,C,,S,A\n,,,S,B\n'   , 'encode index conflict 30'))
T.Pass(encode('[{"key":"K","context":"C","source":"S","values":{"foo":"A"}},{"source":"S","values":{"foo":"B"}}]'                         , 'Key,Context,Example,Source,foo\nK,C,,S,A\n,,,S,B\n'  , 'encode index conflict 31'))
T.Pass(encode('[{"values":{"foo":"A"}},{"key":"K","context":"C","values":{"foo":"B"}}]'                                                   , ERROR_FUNC                                            , 'encode index conflict 32'))
T.Pass(encode('[{"key":"K","values":{"foo":"A"}},{"key":"K","context":"C","values":{"foo":"B"}}]'                                         , ERROR_FUNC                                            , 'encode index conflict 33'))
T.Pass(encode('[{"context":"C","values":{"foo":"A"}},{"key":"K","context":"C","values":{"foo":"B"}}]'                                     , ERROR_FUNC                                            , 'encode index conflict 34'))
T.Pass(encode('[{"source":"S","values":{"foo":"A"}},{"key":"K","context":"C","values":{"foo":"B"}}]'                                      , 'Key,Context,Example,Source,foo\n,,,S,A\nK,C,,,B\n'   , 'encode index conflict 35'))
T.Pass(encode('[{"key":"K","context":"C","values":{"foo":"A"}},{"key":"K","context":"C","values":{"foo":"B"}}]'                           , ERROR_FUNC                                            , 'encode index conflict 36'))
T.Pass(encode('[{"key":"K","source":"S","values":{"foo":"A"}},{"key":"K","context":"C","values":{"foo":"B"}}]'                            , ERROR_FUNC                                            , 'encode index conflict 37'))
T.Pass(encode('[{"context":"C","source":"S","values":{"foo":"A"}},{"key":"K","context":"C","values":{"foo":"B"}}]'                        , 'Key,Context,Example,Source,foo\n,C,,S,A\nK,C,,,B\n'  , 'encode index conflict 38'))
T.Pass(encode('[{"key":"K","context":"C","source":"S","values":{"foo":"A"}},{"key":"K","context":"C","values":{"foo":"B"}}]'              , ERROR_FUNC                                            , 'encode index conflict 39'))
T.Pass(encode('[{"values":{"foo":"A"}},{"key":"K","source":"S","values":{"foo":"B"}}]'                                                    , ERROR_FUNC                                            , 'encode index conflict 40'))
T.Pass(encode('[{"key":"K","values":{"foo":"A"}},{"key":"K","source":"S","values":{"foo":"B"}}]'                                          , ERROR_FUNC                                            , 'encode index conflict 41'))
T.Pass(encode('[{"context":"C","values":{"foo":"A"}},{"key":"K","source":"S","values":{"foo":"B"}}]'                                      , ERROR_FUNC                                            , 'encode index conflict 42'))
T.Pass(encode('[{"source":"S","values":{"foo":"A"}},{"key":"K","source":"S","values":{"foo":"B"}}]'                                       , 'Key,Context,Example,Source,foo\n,,,S,A\nK,,,S,B\n'   , 'encode index conflict 43'))
T.Pass(encode('[{"key":"K","context":"C","values":{"foo":"A"}},{"key":"K","source":"S","values":{"foo":"B"}}]'                            , ERROR_FUNC                                            , 'encode index conflict 44'))
T.Pass(encode('[{"key":"K","source":"S","values":{"foo":"A"}},{"key":"K","source":"S","values":{"foo":"B"}}]'                             , ERROR_FUNC                                            , 'encode index conflict 45'))
T.Pass(encode('[{"context":"C","source":"S","values":{"foo":"A"}},{"key":"K","source":"S","values":{"foo":"B"}}]'                         , 'Key,Context,Example,Source,foo\n,C,,S,A\nK,,,S,B\n'  , 'encode index conflict 46'))
T.Pass(encode('[{"key":"K","context":"C","source":"S","values":{"foo":"A"}},{"key":"K","source":"S","values":{"foo":"B"}}]'               , ERROR_FUNC                                            , 'encode index conflict 47'))
T.Pass(encode('[{"values":{"foo":"A"}},{"context":"C","source":"S","values":{"foo":"B"}}]'                                                , ERROR_FUNC                                            , 'encode index conflict 48'))
T.Pass(encode('[{"key":"K","values":{"foo":"A"}},{"context":"C","source":"S","values":{"foo":"B"}}]'                                      , 'Key,Context,Example,Source,foo\nK,,,,A\n,C,,S,B\n'   , 'encode index conflict 49'))
T.Pass(encode('[{"context":"C","values":{"foo":"A"}},{"context":"C","source":"S","values":{"foo":"B"}}]'                                  , ERROR_FUNC                                            , 'encode index conflict 50'))
T.Pass(encode('[{"source":"S","values":{"foo":"A"}},{"context":"C","source":"S","values":{"foo":"B"}}]'                                   , 'Key,Context,Example,Source,foo\n,,,S,A\n,C,,S,B\n'   , 'encode index conflict 51'))
T.Pass(encode('[{"key":"K","context":"C","values":{"foo":"A"}},{"context":"C","source":"S","values":{"foo":"B"}}]'                        , 'Key,Context,Example,Source,foo\nK,C,,,A\n,C,,S,B\n'  , 'encode index conflict 52'))
T.Pass(encode('[{"key":"K","source":"S","values":{"foo":"A"}},{"context":"C","source":"S","values":{"foo":"B"}}]'                         , 'Key,Context,Example,Source,foo\nK,,,S,A\n,C,,S,B\n'  , 'encode index conflict 53'))
T.Pass(encode('[{"context":"C","source":"S","values":{"foo":"A"}},{"context":"C","source":"S","values":{"foo":"B"}}]'                     , ERROR_FUNC                                            , 'encode index conflict 54'))
T.Pass(encode('[{"key":"K","context":"C","source":"S","values":{"foo":"A"}},{"context":"C","source":"S","values":{"foo":"B"}}]'           , 'Key,Context,Example,Source,foo\nK,C,,S,A\n,C,,S,B\n' , 'encode index conflict 55'))
T.Pass(encode('[{"values":{"foo":"A"}},{"key":"K","context":"C","source":"S","values":{"foo":"B"}}]'                                      , ERROR_FUNC                                            , 'encode index conflict 56'))
T.Pass(encode('[{"key":"K","values":{"foo":"A"}},{"key":"K","context":"C","source":"S","values":{"foo":"B"}}]'                            , ERROR_FUNC                                            , 'encode index conflict 57'))
T.Pass(encode('[{"context":"C","values":{"foo":"A"}},{"key":"K","context":"C","source":"S","values":{"foo":"B"}}]'                        , ERROR_FUNC                                            , 'encode index conflict 58'))
T.Pass(encode('[{"source":"S","values":{"foo":"A"}},{"key":"K","context":"C","source":"S","values":{"foo":"B"}}]'                         , 'Key,Context,Example,Source,foo\n,,,S,A\nK,C,,S,B\n'  , 'encode index conflict 59'))
T.Pass(encode('[{"key":"K","context":"C","values":{"foo":"A"}},{"key":"K","context":"C","source":"S","values":{"foo":"B"}}]'              , ERROR_FUNC                                            , 'encode index conflict 60'))
T.Pass(encode('[{"key":"K","source":"S","values":{"foo":"A"}},{"key":"K","context":"C","source":"S","values":{"foo":"B"}}]'               , ERROR_FUNC                                            , 'encode index conflict 61'))
T.Pass(encode('[{"context":"C","source":"S","values":{"foo":"A"}},{"key":"K","context":"C","source":"S","values":{"foo":"B"}}]'           , 'Key,Context,Example,Source,foo\n,C,,S,A\nK,C,,S,B\n' , 'encode index conflict 62'))
T.Pass(encode('[{"key":"K","context":"C","source":"S","values":{"foo":"A"}},{"key":"K","context":"C","source":"S","values":{"foo":"B"}}]' , ERROR_FUNC                                            , 'encode index conflict 63'))

-- Test merged headers.
T.Pass(encode('[{"key":"11","values":{"A":"AA","B":"BB"}},{"key":"22","values":{"B":"BB","C":"CC"}}]', 'Key,Context,Example,Source,a,b,c\n11,,,,AA,BB,\n22,,,,,BB,CC\n', 'encode merge headers'))

-- Test CSV format.
T.Pass(encode('[{"key":"KK","values":{"\\"g\\"":"GG","a":"AA","b":"BB","c d":"CD","e f":"EF"}}]', 'Key,Context,Example,Source,"""g""",a,b,c d,e f\nKK,,,,GG,AA,BB,CD,EF\n', 'encode csv format'))
