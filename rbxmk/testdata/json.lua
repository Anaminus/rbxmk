T.Pass(json.fromString("null") == nil, "null from string")
T.Pass(select("#", json.fromString("null")) > 0, "null from string is not empty")
T.Pass(json.fromString("false") == false, "false from string")
T.Pass(json.fromString("true") == true, "true from string")
T.Pass(json.fromString("3.14159") == 3.14159, "3.14159 from string")
T.Pass(json.fromString("-3.14159") == -3.14159, "-3.14159 from string")
T.Pass(json.fromString("\"foobar\"") == "foobar", "\"foobar\" from string")

local array = json.fromString("[1,2,3]")
T.Pass(type(array) == "table", "array from string is a table")
T.Pass(#array == 3, "array from string has 3 elements")
T.Pass(array[1] == 1, "array[1] from string is 1")
T.Pass(array[2] == 2, "array[2] from string is 2")
T.Pass(array[3] == 3, "array[3] from string is 3")

local object = json.fromString([[{"foo": 1, "bar": 2}]])
T.Pass(type(object) == "table", "object from string is a table")
T.Pass(#object == 0, "object from string has 0 elements")
T.Pass(object.foo == 1, "object.foo from string is 1")
T.Pass(object.bar == 2, "object.bar from string is 2")

T.Pass(json.string(array) == "[\n\t1,\n\t2,\n\t3\n]", "array to string with tabs")
T.Pass(json.string(object) == '{\n\t"bar": 2,\n\t"foo": 1\n}', "object to string with tabs")

T.Pass(json.string(array,"  ") == "[\n  1,\n  2,\n  3\n]", "array to string with spaces")
T.Pass(json.string(object,"  ") == '{\n  "bar": 2,\n  "foo": 1\n}', "object to string with spaces")

T.Pass(json.string(array,"") == "[1,2,3]", "array to string minified")
T.Pass(json.string(object,"") == '{"bar":2,"foo":1}', "object to string minified")

local big = [[{"a":[1,2,3],"b":[4,5,6],"c":{"x":true,"y":false,"z":42}}]]
T.Pass(json.string(json.fromString(big),"") == big, "roundtrip")

local prev = {foo=1,bar=2}
local next = {foo=1,bar=3,baz=2}
local patch = json.diff(prev,next)
local result = json.patch(prev, patch)

T.Pass(type(patch) == "table", "patch is a table")
T.Pass(#patch == 2, "patch has 2 elements")
T.Pass(patch[1].op == "replace", "patch[1].op == replace")
T.Pass(patch[1].path == "/bar", "patch[1].path == /bar")
T.Pass(patch[1].value == 3, "patch[1].value == 3")
T.Pass(patch[2].op == "add", "patch[2].op == add")
T.Pass(patch[2].path == "/baz", "patch[2].path == /baz")
T.Pass(patch[2].value == 2, "patch[2].value == 2")

T.Pass(json.string(result,"") == json.string(next,""), "result equals next")
