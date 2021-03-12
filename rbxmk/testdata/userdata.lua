T.GC()
local baseLen = T.UserDataCacheLen()

-- Test cached userdata.
local game = DataModel.new()
local object0 = Instance.new("Object", game)
T.Pass(T.UserDataCacheLen() == baseLen+2, "userdata cache length is 2")

local set = {[object0] = true}
local object1 = game:GetChildren()[1]
T.Pass(set[object1] == true, "cached userdata is returned")

game = nil
T.GC()
T.Pass(T.UserDataCacheLen() == baseLen+1, "userdata cache length is 1")

object0 = nil
T.GC()
T.Pass(T.UserDataCacheLen() == baseLen+1, "userdata cache length is 1")

object1 = nil
T.GC()
T.Pass(T.UserDataCacheLen() == baseLen+1, "userdata cache length is 1")

set = nil
T.GC()
-- Requires two cycles to fully finalize.
T.GC()
T.Pass(T.UserDataCacheLen() == baseLen+0, "userdata cache length is 0")

T.GC()

-- Test uncached userdata.
local v0 = Vector3.new(1,2,3)
local v1 = Vector3.new(1,2,3)
local v2 = Vector3.new(4,5,6)
T.Pass(T.UserDataCacheLen() == baseLen, "userdata cache length is 0")
T.Pass(v0 == v0, "v0 == v0")
T.Pass(v0 == v1, "v0 == v1")
T.Pass(v0 ~= v2, "v0 ~= v2")
T.Pass(v1 == v0, "v1 == v0")
T.Pass(v1 == v1, "v1 == v1")
T.Pass(v1 ~= v2, "v1 ~= v2")
T.Pass(v2 ~= v0, "v2 ~= v0")
T.Pass(v2 ~= v1, "v2 ~= v1")
T.Pass(v2 == v2, "v2 == v2")

local set = {}
set[v0] = 0
T.Pass(set[v0] == 0, "v0 is returned")
T.Pass(set[v1] == nil, "v1 is not returned")
T.Pass(set[v2] == nil, "v2 is not returned")
set[v1] = 1
T.Pass(set[v0] == 0, "v0 is returned")
T.Pass(set[v1] == 1, "v1 is returned")
T.Pass(set[v2] == nil, "v2 is not returned")
set[v2] = 2
T.Pass(set[v0] == 0, "v0 is returned")
T.Pass(set[v1] == 1, "v1 is returned")
T.Pass(set[v2] == 2, "v2 is returned")
