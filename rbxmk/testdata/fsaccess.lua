T.Fail(function() return fs.stat("/") ~= nil end, "fs root")

T.Fail(function() return fs.stat(os.expand("$wd")) ~= nil end, "stat working directory")
T.Fail(function() return fs.stat(".") ~= nil end, "stat working directory w/ dot")
T.Pass(function() return fs.dir(".") ~= nil end, "dir working directory")

T.Pass(function() return fs.stat(os.expand("$sd")) ~= nil end, "script directory")

T.Pass(function() return fs.stat(os.expand("$rsd")) ~= nil end, "root script directory")

T.Pass(function() return fs.stat(os.expand("$wd/foo/bar/baz")) == nil end, "nonextant")

T.Fail(function() return fs.stat("..") == nil end, "wd parent")
T.Fail(function() return fs.stat("../..") == nil end, "wd ancestor")
T.Fail(function() return fs.stat("../../foo/bar") == nil end, "out then elsewhere")
T.Pass(function() return fs.stat("../../rbxmk/rbxmk/foo") == nil end, "out then back in")
