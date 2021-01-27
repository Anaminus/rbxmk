T.Pass(rbxmk.decodeFormat("json", [[null]])            == nil             , "json null decodes into nil")
T.Pass(rbxmk.decodeFormat("json", [[false]])           == false           , "json false decodes into false")
T.Pass(rbxmk.decodeFormat("json", [[true]])            == true            , "json true decodes into true")
T.Pass(rbxmk.decodeFormat("json", [[-3.14159]])        == -3.14159        , "json number decodes into number")
T.Pass(rbxmk.decodeFormat("json", [["Hello, world!"]]) == "Hello, world!" , "json string decodes into string")

local array = rbxmk.decodeFormat("json", [[ [1,2,3] ]])
T.Pass(type(array) == "table" , "json array decodes into table")
T.Pass(array[1] == 1          , "decodes array index 1")
T.Pass(array[2] == 2          , "decodes array index 2")
T.Pass(array[3] == 3          , "decodes array index 3")

local object = rbxmk.decodeFormat("json", [[ {"a":1,"b":2,"c":3} ]])
T.Pass(type(object) == "table" , "json object decodes into table")
T.Pass(object.a == 1           , "decodes object field a")
T.Pass(object.b == 2           , "decodes object field b")
T.Pass(object.c == 3           , "decodes object field c")

T.Pass(rbxmk.encodeFormat("json", nil)             == "null\n"            , "nil encodes into json null")
T.Pass(rbxmk.encodeFormat("json", false)           == "false\n"           , "false encodes into json false")
T.Pass(rbxmk.encodeFormat("json", true)            == "true\n"            , "true encodes into json true")
T.Pass(rbxmk.encodeFormat("json", -3.14159)        == "-3.14159\n"        , "number encodes into json number")
T.Pass(rbxmk.encodeFormat("json", "Hello, world!") == '"Hello, world!"\n' , "string encodes into json string")

local jarray = '[\n\t1,\n\t2,\n\t3\n]\n'
T.Pass(rbxmk.encodeFormat("json", {1,2,3}) == jarray, "array table encodes into json array")

local jobject = '{\n\t"a": 1,\n\t"b": 2,\n\t"c": 3\n}\n'
T.Pass(rbxmk.encodeFormat("json", {a=1,b=2,c=3}) == jobject, "dictionary table encodes into json object")
