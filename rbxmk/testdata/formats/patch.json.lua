T.Pass(rbxmk.decodeFormat("patch.json", [[ [] ]]), "empty patch")
T.Pass(rbxmk.decodeFormat("patch.json", [[ [{"op":"test"}] ]])[1].op == "test", "test operation")
T.Pass(rbxmk.decodeFormat("patch.json", [[ [{"op":"remove"}] ]])[1].op == "remove", "remove operation")
T.Pass(rbxmk.decodeFormat("patch.json", [[ [{"op":"add"}] ]])[1].op == "add", "add operation")
T.Pass(rbxmk.decodeFormat("patch.json", [[ [{"op":"replace"}] ]])[1].op == "replace", "replace operation")
T.Pass(rbxmk.decodeFormat("patch.json", [[ [{"op":"move"}] ]])[1].op == "move", "move operation")
T.Pass(rbxmk.decodeFormat("patch.json", [[ [{"op":"copy"}] ]])[1].op == "copy", "copy operation")
T.Fail(function() rbxmk.decodeFormat("patch.json", [[ [{"op":"foo"}] ]]) end, "foo operation")

local big = [[ [
	{"op": "test"    , "from": "/foo", "path": "/bar", "value": [1,2,3]},
	{"op": "remove"  , "from": "/foo", "path": "/bar", "value": [1,2,3]},
	{"op": "add"     , "from": "/foo", "path": "/bar", "value": [1,2,3]},
	{"op": "replace" , "from": "/foo", "path": "/bar", "value": [1,2,3]},
	{"op": "move"    , "from": "/foo", "path": "/bar", "value": [1,2,3]},
	{"op": "copy"    , "from": "/foo", "path": "/bar", "value": [1,2,3]}
] ]]

local patch = rbxmk.decodeFormat("patch.json", big)
T.Pass(type(patch) == "table", "patch is table")
T.Pass(#patch == 6, "patch contains 6 operations")

T.Pass(patch[1].op == "test", "patch[1].op is test")
T.Pass(patch[1].path == "/bar", "patch[1].path is /bar")
T.Pass(patch[1].from == nil, "patch[1].from is nil; not applicable for op")
T.Pass(#patch[1].value == 3, "patch[1].value is array")

T.Pass(patch[2].op == "remove", "patch[2].op is remove")
T.Pass(patch[2].path == "/bar", "patch[2].path is /bar")
T.Pass(patch[2].from == nil, "patch[2].from is nil; not applicable for op")
T.Pass(patch[2].value == nil, "patch[2].value is nil; not applicable for op")

T.Pass(patch[3].op == "add", "patch[3].op is add")
T.Pass(patch[3].path == "/bar", "patch[3].path is /bar")
T.Pass(patch[3].from == nil, "patch[3].from is nil; not applicable for op")
T.Pass(#patch[3].value == 3, "patch[3].value is array")

T.Pass(patch[4].op == "replace", "patch[4].op is replace")
T.Pass(patch[4].path == "/bar", "patch[4].path is /bar")
T.Pass(patch[4].from == nil, "patch[4].from is nil; not applicable for op")
T.Pass(#patch[4].value == 3, "patch[4].value is array")

T.Pass(patch[5].op == "move", "patch[5].op is move")
T.Pass(patch[5].path == "/bar", "patch[5].path is /bar")
T.Pass(patch[5].from == "/foo", "patch[5].from is /foo")
T.Pass(patch[5].value == nil, "patch[5].value is nil; not applicable for op")

T.Pass(patch[6].op == "copy", "patch[6].op is copy")
T.Pass(patch[6].path == "/bar", "patch[6].path is /bar")
T.Pass(patch[6].from == "/foo", "patch[6].from is /foo")
T.Pass(patch[6].value == nil, "patch[6].value is nil; not applicable for op")
