T.GC()
local baseLen = T.UserDataCacheLen()

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
T.Pass(T.UserDataCacheLen() == baseLen+0, "userdata cache length is 0")
