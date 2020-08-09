T.GC()
local baseLen = T.UserDataCacheLen()

local game = DataModel.new()
local object0 = Instance.new("Object", game)
T.Pass("userdata cache length is 2", T.UserDataCacheLen() == baseLen+2)

local set = {[object0] = true}
local object1 = game:GetChildren()[1]
T.Pass("cached userdata is returned", set[object1] == true)

game = nil
T.GC()
T.Pass("userdata cache length is 1", T.UserDataCacheLen() == baseLen+1)

object0 = nil
T.GC()
T.Pass("userdata cache length is 1", T.UserDataCacheLen() == baseLen+1)

object1 = nil
T.GC()
T.Pass("userdata cache length is 1", T.UserDataCacheLen() == baseLen+1)

set = nil
T.GC()
T.Pass("userdata cache length is 0", T.UserDataCacheLen() == baseLen+0)
