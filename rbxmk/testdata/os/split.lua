local path = "project/scripts/main.script.lua"

local c = table.pack(os.split(path))
T.Pass(c.n == 0, "length of none")

local c = table.pack(os.split(path, "base"))
T.Pass(c.n == 1, "length of base")
T.Pass(c[1] == "main.script.lua", "base")

local c = table.pack(os.split(path, "dir"))
T.Pass(c.n == 1, "length of dir")
T.Pass(c[1] == os.clean("project/scripts"), "dir")

local c = table.pack(os.split(path, "ext"))
T.Pass(c.n == 1, "length of ext")
T.Pass(c[1] == ".lua", "ext")

local c = table.pack(os.split(path, "fext"))
T.Pass(c.n == 1, "length of fext")
T.Pass(c[1] == ".script.lua", "fext")

local c = table.pack(os.split(path, "fstem"))
T.Pass(c.n == 1, "length of fstem")
T.Pass(c[1] == "main", "fstem")

local c = table.pack(os.split(path, "stem"))
T.Pass(c.n == 1, "length of stem")
T.Pass(c[1] == "main.script", "stem")

local c = table.pack(os.split(path, "dir", "base", "ext", "fext", "stem", "fstem"))
T.Pass(c.n == 6, "length of all")
T.Pass(c[1] == os.clean("project/scripts"), "all dir")
T.Pass(c[2] == "main.script.lua", "all base")
T.Pass(c[3] == ".lua", "all ext")
T.Pass(c[4] == ".script.lua", "all fext")
T.Pass(c[5] == "main.script", "all stem")
T.Pass(c[6] == "main", "all fstem")

T.Fail(function() return os.split(path, "UNKNOWN") end, "unknown component")
